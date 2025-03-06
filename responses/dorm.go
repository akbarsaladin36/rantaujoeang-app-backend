package responses

import (
	"rantaujoeang-app-backend/models"
	"time"
)

type DormResponse struct {
	DormCode           string `json:"dorm_code"`
	DormName           string `json:"dorm_name"`
	DormPriceAmount    string `json:"dorm_price_amount"`
	DormQuantityAmount string `json:"dorm_quantity_amount"`
	DormAddress        string `json:"dorm_address"`
	DormDescription    string `json:"dorm_description"`
	DormPhoneNumber    string `json:"dorm_phone_number"`
}

type createDormResponse struct {
	DormCode                string `json:"dorm_code"`
	DormName                string `json:"dorm_name"`
	DormPriceAmount         string `json:"dorm_price_amount"`
	DormQuantityAmount      string `json:"dorm_quantity_amount"`
	DormAddress             string `json:"dorm_address"`
	DormDescription         string `json:"dorm_description"`
	DormPhoneNumber         string `json:"dorm_phone_number"`
	DormCreatedDate         time.Time
	DormCreatedUserUuid     string `json:"dorm_created_user_uuid"`
	DormCreatedUserUsername string `json:"dorm_created_user_username"`
}

type updateDormResponse struct {
	DormCode                string `json:"dorm_code"`
	DormName                string `json:"dorm_name"`
	DormPriceAmount         string `json:"dorm_price_amount"`
	DormQuantityAmount      string `json:"dorm_quantity_amount"`
	DormAddress             string `json:"dorm_address"`
	DormDescription         string `json:"dorm_description"`
	DormPhoneNumber         string `json:"dorm_phone_number"`
	DormUpdatedDate         time.Time
	DormUpdatedUserUuid     string `json:"dorm_updated_user_uuid"`
	DormUpdatedUserUsername string `json:"dorm_updated_user_username"`
}

func GetDormResponse(dormRsps models.Dorm) DormResponse {
	return DormResponse{
		DormCode:           dormRsps.DormCode,
		DormName:           dormRsps.DormName,
		DormPriceAmount:    dormRsps.DormPriceAmount,
		DormQuantityAmount: dormRsps.DormQuantityAmount,
		DormAddress:        dormRsps.DormAddress,
		DormDescription:    dormRsps.DormDescription,
		DormPhoneNumber:    dormRsps.DormPhoneNumber,
	}
}

func GetCreateDormResponse(dormRsps models.Dorm) createDormResponse {
	return createDormResponse{
		DormCode:                dormRsps.DormCode,
		DormName:                dormRsps.DormName,
		DormPriceAmount:         dormRsps.DormPriceAmount,
		DormQuantityAmount:      dormRsps.DormQuantityAmount,
		DormAddress:             dormRsps.DormAddress,
		DormDescription:         dormRsps.DormDescription,
		DormPhoneNumber:         dormRsps.DormPhoneNumber,
		DormCreatedDate:         dormRsps.DormCreatedDate,
		DormCreatedUserUuid:     dormRsps.DormCreatedUserUuid,
		DormCreatedUserUsername: dormRsps.DormCreatedUserUsername,
	}
}

func GetUpdateDormResponse(dormRsps models.Dorm) updateDormResponse {
	return updateDormResponse{
		DormCode:                dormRsps.DormCode,
		DormName:                dormRsps.DormName,
		DormPriceAmount:         dormRsps.DormPriceAmount,
		DormQuantityAmount:      dormRsps.DormQuantityAmount,
		DormAddress:             dormRsps.DormAddress,
		DormDescription:         dormRsps.DormDescription,
		DormPhoneNumber:         dormRsps.DormPhoneNumber,
		DormUpdatedDate:         dormRsps.DormUpdatedDate,
		DormUpdatedUserUuid:     dormRsps.DormUpdatedUserUuid,
		DormUpdatedUserUsername: dormRsps.DormUpdatedUserUsername,
	}
}
