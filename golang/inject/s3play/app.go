package s3play

import "github.com/katsuokaisao/localstack/usecase/s3play"

type Application struct {
	S3PlayUseCase s3play.S3PlayUseCase
}

func NewApplication(
	s3PlayUseCase s3play.S3PlayUseCase,
) *Application {
	return &Application{
		S3PlayUseCase: s3PlayUseCase,
	}
}

func (a *Application) Run() {
	a.S3PlayUseCase.Play()
}