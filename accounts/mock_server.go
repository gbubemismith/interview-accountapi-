package accounts

var (
	mockupServer = mockServer{
		mocks: make(map[string]*Mock),
	}
)

type mockServer struct {
	enabled bool

	mocks map[string]*Mock
}

func StartMockServer() {
	mockupServer.enabled = true
}

func StopMockServer() {
	mockupServer.enabled = false
}

func AddMock(mock Mock) {
	key := mock.Method + mock.Url + mock.RequestBody
	mockupServer.mocks[key] = &mock
}
