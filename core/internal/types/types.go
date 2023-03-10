// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserDetailRequest struct {
	Identity int64 `json:"identity"`
}

type UserDetailResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MailCodeSendRequest struct {
	Email string `json:"email"`
}

type MailCodeSendResponse struct {
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type UserRegisterResponse struct {
	UserId  string `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Path string `json:"path,optional"`
	Size int64  `json:"size,optional"`
}

type FileUploadResponse struct {
	Identity int64  `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type UserRepositoryRequest struct {
	ParentId           int    `json:"parentId"`
	RepositoryIdentity int64  `json:"repository_identity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositoryResponse struct {
	Identity int64 `json:"identity"`
}

type UserFileListRequest struct {
	Id   int `json:"id,optional"`
	Page int `json:"page,optional"`
	Size int `json:"size,optional"`
}

type UserFile struct {
	Id                 int    `json:"id"`
	Identity           int64  `json:"identity"`
	RepositoryIdentity int64  `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	Size               int64  `json:"size"`
}

type UserFileListResponse struct {
	List  []*UserFile `json:"list"`
	Count int64       `json:"count"`
}

type UserFileNameModifyRequest struct {
	Identity int64  `json:"identity"`
	Name     string `json:"name"`
}

type UserFileNameModifyResponse struct {
}

type UserFolderCreateRequest struct {
	ParentId int    `json:"parent_id"`
	Name     string `json:"name"`
}

type UserFolderCreateResponse struct {
	Identity int64 `json:"identity"`
}

type UserFileDeleteRequest struct {
	Identity int64 `json:"identity"`
}

type UserFileDeleteResponse struct {
}

type UserFileMoveRequest struct {
	Identity       int64 `json:"identity"`
	ParentIdentity int64 `json:"parent_identity"`
}

type UserFileMoveResponse struct {
}

type ShareBasicCreateRequest struct {
	UserRepositoryIdentity int64 `json:"user_repository_identity"`
	ExpiredTime            int   `json:"expired_time"`
}

type ShareBasicCreateResponse struct {
	Identity int64 `json:"identity"`
}

type ShareBasicDetailRequest struct {
	Identity int64 `json:"identity"`
}

type ShareBasicDetailResponse struct {
	RepositoryIdentity int64  `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Size               int64  `json:"size"`
	Path               string `json:"path"`
}

type ShareBasicSaveRequest struct {
	RepositoryIdentity int64 `json:"repository_identity"`
	ParentId           int   `json:"parent_id"`
}

type ShareBasicSaveResponse struct {
	Identity int64 `json:"identity"`
}

type RefreshAuthorizationRequest struct {
}

type RefreshAuthorizationResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type FileUploadPrepareRequest struct {
	Md5  string `json:"md5"`
	Name string `json:"name"`
}

type FileUploadPrepareResponse struct {
	Identity int64 `json:"identity"`
	UploadId int   `json:"upload_id"`
}
