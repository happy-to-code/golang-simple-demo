package example

// repo

// IPostRepo IPostRepo
type IPostRepo interface{}

// NewPostRepo NewPostRepo
func NewPostRepo() IPostRepo {
	return new(IPostRepo)
}

// usecase

// IPostUsecase IPostUsecase
type IPostUsecase interface{}
type postUsecase struct {
	repo IPostRepo
}

// NewPostUsecase NewPostUsecase Provider  函数
func NewPostUsecase(repo IPostRepo) IPostUsecase {
	return postUsecase{repo: repo}
}

// service service

// PostService PostService
type PostService struct {
	usecase IPostUsecase
}

// NewPostService NewPostService  Provider  函数
func NewPostService(u IPostUsecase) *PostService {
	return &PostService{usecase: u}
}
