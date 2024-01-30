# Unit Test
- Follow the structure as mentioned in user_service_test.go file to create test cases
- Structure for the test will look like this
 ```
tests := []struct {
		name              string     // name of test case
		args              test.Args  // the stuct contains the request
		expectedResponse  test.Response // expected response
		expectedError     error     // expected error
		mock              func()    // External calls inside the function will be mocked here
		tests             []string  // The comparison that needs to be done on the test case
	}
 ```
- The request and response for the test cases will be stores in tests folder. Inside tests folder we will have service folder.
- Inside the service folder, for each function we will have a .go file, where the request, response and error will be stored mapped with case
- Check the Test_validate_user.go file for the same

## Mock
- Using testify we will mock the outside calls which are within the function
 ```
mock: func() {
				meetingInfo := &pb.MeetingInfo{
					Id: "1",
					Topics: []*pb.Topic{
						{
							Id: "1",
						},
					},
					Attendees: []*pb.Attendee{
						{
							Id:   "1",
							Type: pb.AttendeeType_TEACHER,
							Role: pb.AttendeeRole_ORGANISER,
						},
						{
							Id:   "2",
							Type: pb.AttendeeType_BATCH,
							Role: pb.AttendeeRole_AUDIENCE,
						},
					},
				}
				mockHandler.On("CreateMeeting", ctx, meetingInfo).Return(meetingInfo, nil)
				mockTopicService.On("createTopics", ctx, meetingInfo.Id, meetingInfo.Topics).Return(nil)
				mockAttendeesService.On("addAttendees", ctx, meetingInfo.Id, meetingInfo.Attendees).Return(nil)
			}	
 ```