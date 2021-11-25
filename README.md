# Form3 Take Home Exercise

### Author -> Gbubemi Smith
I am new to golang(3 weeks+) and this is my first project. I tried to build the library following all the principles I have become accustomed to as a developer e.g OOP, SOLID. I ensured any user of the library only has access to the public interfaces which I exposed and I tried as much as possible to encapsulate and hide much of the acutal implementation. There are two main folders in this project namely: form3htttp, accounts. # form3http is basically a wrapper around the required http calls to the form3 api.I used a builder approach as well exposing functionalities through abstarct methods rather than concrete implementattions. The form3http is resuable and all that needs to be passed basically are the headers, api url and body. # accounts is the actual client library for this project and I made it as simple as possible to initialize and call. An Init function is used to intialize the library and the base address is an argument that needs to be passed. Once that is done the available functionalities based on this project # shoulds would be available to be called. The test cases are also written in this directory and can be seen once the command docker-compose up is entered. There is an example folder where I have example calls. Some tutorials and articles where referenced to make this happen
- https://blog.logrocket.com/making-http-requests-in-go
- https://www.digitalocean.com/community/tutorials/defining-methods-in-go
- https://tour.golang.org/
- https://pkg.go.dev/github.com/stretchr/testify/assert
- https://devcharmander.medium.com/design-patterns-in-golang-the-builder-dac468a71194
- https://www.youtube.com/watch?v=uB_45bSIyik&t=972s

Engineers at Form3 build highly available distributed systems in a microservices environment. Our take home test is designed to evaluate real world activities that are involved with this role. We recognise that this may not be as mentally challenging and may take longer to implement than some algorithmic tests that are often seen in interview exercises. Our approach however helps ensure that you will be working with a team of engineers with the necessary practical skills for the role (as well as a diverse range of technical wizardry). 

## Instructions
The goal of this exercise is to write a client library in Go to access our fake account API, which is provided as a Docker
container in the file `docker-compose.yaml` of this repository. Please refer to the
[Form3 documentation](http://api-docs.form3.tech/api.html#organisation-accounts) for information on how to interact with the API. Please note that the fake account API does not require any authorisation or authentication.

A mapping of account attributes can be found in [models.go](./models.go). Can be used as a starting point, usage of the file is not required.

If you encounter any problems running the fake account API we would encourage you to do some debugging first,
before reaching out for help.

## Submission Guidance

### Shoulds

The finished solution **should:**
- Be written in Go.
- Use the `docker-compose.yaml` of this repository.
- Be a client library suitable for use in another software project.
- Implement the `Create`, `Fetch`, and `Delete` operations on the `accounts` resource.
- Be well tested to the level you would expect in a commercial environment. Note that tests are expected to run against the provided fake account API.
- Be simple and concise.
- Have tests that run from `docker-compose up` - our reviewers will run `docker-compose up` to assess if your tests pass.

### Should Nots

The finished solution **should not:**
- Use a code generator to write the client library.
- Use (copy or otherwise) code from any third party without attribution to complete the exercise, as this will result in the test being rejected.
- Use a library for your client (e.g: go-resty). Libraries to support testing or types like UUID are fine.
- Implement client-side validation.
- Implement an authentication scheme.
- Implement support for the fields `data.attributes.private_identification`, `data.attributes.organisation_identification`
  and `data.relationships`, as they are omitted in the provided fake account API implementation.
- Have advanced features, however discussion of anything extra you'd expect a production client to contain would be useful in the documentation.
- Be a command line client or other type of program - the requirement is to write a client library.
- Implement the `List` operation.
> We give no credit for including any of the above in a submitted test, so please only focus on the "Shoulds" above.

## How to submit your exercise

- Include your name in the README. If you are new to Go, please also mention this in the README so that we can consider this when reviewing your exercise
- Create a private [GitHub](https://help.github.com/en/articles/create-a-repo) repository, by copying all files you deem necessary for your submission
- [Invite](https://help.github.com/en/articles/inviting-collaborators-to-a-personal-repository) @form3tech-interviewer-1 to your private repo
- Let us know you've completed the exercise using the link provided at the bottom of the email from our recruitment team

## License

Copyright 2019-2021 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
