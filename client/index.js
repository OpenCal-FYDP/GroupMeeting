import { client } from "twirpscript";
import { nodeHttpTransport } from "twirpscript/dist/node/index.js";
import {GetGroupEventJSON, UpdateGroupEvent, AttendeeAvailability, GetGroupEventReq, UpdateGroupEventReq, GetGroupEventRes} from "./service.pb.js";

client.baseURL = "http://localhost:8080";

// This is provided as a convenience for Node.js clients. If you provide `fetch` globally, this isn't necessary and your client can look identical to the browser client above.
client.rpcTransport = nodeHttpTransport;


const test1AA = AttendeeAvailability.initialize();
test1AA.availabilityID = "3";
test1AA.DateRanges = ["11111-11112", "2222-2223", "3333-3334"];

const test = UpdateGroupEventReq.initialize();
console.log(test);
test.attendees = ["test1@gmail.com", "test2@gmail.com"];
test.availabilities = {"test1@gmail.com" : test1AA};
test.eventID = "7";
console.log(test);

const update = await UpdateGroupEvent(test);

var test2 = GetGroupEventReq.initialize();
test2.eventID = "7";
const get = await GetGroupEventJSON(test2);
// const profile = await SetUserProfile({
//     userID: "test-user-id",
//     email: "test@test.com",
//     username: "test name",
// } );

// const profile = await GetUserProfile({
//     userID: "test-user-id",
// } );

console.log(get);




