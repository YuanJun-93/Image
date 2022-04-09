package mysql

func GetImageByID(imageID int64) (path string, err error){
	sqlStr := `select image_url from store where image_id = ?`
	err =  db.Get(&path, sqlStr, imageID)
	return path, err
}
