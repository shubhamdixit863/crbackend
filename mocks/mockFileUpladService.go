package mocks

type MockFileUploadService struct {
}

func (m *MockFileUploadService) UploadMultipleFile(data []string) ([]string, error) {
	return nil, nil

}