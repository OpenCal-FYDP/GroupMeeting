package storage

import (
	"errors"
	"github.com/OpenCal-FYDP/GroupMeeting/rpc"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"fmt"
)

const tableName = "GroupEvents"

type Storage struct {
	client dynamodbiface.DynamoDBAPI
}

// Stealing Shiv's code, but initializing session to database
func New() *Storage {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	
	client := dynamodb.New(sess)

	return &Storage{client}
}


// Get is just a simple go to the database, and update the response with the values. Annd return
func (s *Storage) GetGroupEvent(req *rpc.GetGroupEventReq, res *rpc.GetGroupEventRes) error {
	result, err := s.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"eventID": {
				S: aws.String(req.EventID),
			},

		},

	})

	if err != nil {
		return err
	}

	// Now we will need to go through and update the res
	if result.Item == nil {
		return errors.New("key not found")
	}

	// Unmarshal the value from dynamodb attribute to go type
	// It will be stored in res address
	err = dynamodbattribute.UnmarshalMap(result.Item, &res)
	if err != nil {
		return err
	}


	return nil
}

func StoreGroupEvent(req *rpc.UpdateGroupEventReq, s *Storage) error {
	av, err := dynamodbattribute.MarshalMap(req)

	if err != nil {
		return err
	}

	// Now create put item
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	// Now we push it into the database, then return
	_, err = s.client.PutItem(input)
	
	return err
}


// Update is different. 
// Update the database with the request.
// If the attendee map no longer has a null for daterange, we will emite an event request to the calender service
func (s *Storage) UpdateGroupEvent(req *rpc.UpdateGroupEventReq, res *rpc.UpdateGroupEventRes) error {
	// First need to get the event. 
	result, err := s.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"eventID": {
				S: aws.String(req.GetEventID()),
			},
			
		},
		
	})
	
	if err != nil {
		return err
	}
	
	// If theres no event, then we will create and yeet out.
	if result.Item == nil {
		return StoreGroupEvent(req, s)
	}
	
	// If there is an event, then I will need to add into the map from database all the values from the req. This includes if someone updated their availibility
	databaseVal := new(rpc.UpdateGroupEventReq)
	err = dynamodbattribute.UnmarshalMap(result.Item, &databaseVal)

	if err != nil {
		return err
	}

	for key, attendeeAvaVal := range req.GetAvailabilities() {
		databaseVal.Availabilities[key] = attendeeAvaVal
	}
	// Once len(attendees) = len(map) after combining, then we emit an event to calender lambda
	if len(req.GetAttendees()) == len(databaseVal.GetAvailabilities()) {
		// sennd a request to create the calender event
		fmt.Println("Sent Request")
	}

	// re-store the values
	return StoreGroupEvent(databaseVal, s)
}

