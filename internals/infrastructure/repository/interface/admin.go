package interfacerepository

type IAdminRepository interface{
	GetPassword( string) (string, error) 
}