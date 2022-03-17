package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"image/model"
)

func UploadImages(p *model.Image) (err error) {
	sqlStr := `insert into store(image_id,user_id,image_cost,image_name,image_url) values(?,?,?,?,?);`
	_, err = db.Exec(sqlStr, p.ImageID, p.UserID, p.ImageCost, p.ImageName, p.ImageUrl)
	return
}

func GetImages() (data []*model.Image, err error) {
	sqlStr := `select image_id, user_id, image_cost, image_name, image_url from store`
	if err = db.Select(&data, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("there is no image info in db")
			err = nil
		}
	}
	return
}
