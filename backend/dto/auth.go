package dto

type LoginDTO struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
<<<<<<< HEAD
	Password string `json:"password" binding:"required,min=6,max=72"`
}

type RegisterDTO struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=6,max=72"`
=======
	Password string `json:"password" binding:"required,min=6"`
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
}