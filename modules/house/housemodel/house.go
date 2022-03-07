package housemodel
import (
  "hfs_backend/common"
  "strings"
)
const EntityName = "House"
type House struct {
	common.SQLModel `json:",inline"`
	OwnerId             int             `json:"owner_id" gorm:"column:owner_id;"`
	Name                string          `json:"name" gorm:"column:name;"`
	Description         string          `json:"description" gorm:"description;"`
	Price               float64         `json:"price" gorm:"price;"`
	lat                 float64         `json:"lat" gorm:"column:lat;"`
	lng                 float64         `json:"lng" gorm:"column:lng;"`
	PhoneNumber         string          `json:"phone_number" gorm:"column:phone_number;"`
  airConditioner      bool            `json:"air_conditioner" gorm:"column:air_conditioner;"`
  RoomBathroom        bool            `json:"room_bathroom" gorm:"column:room_bathroom;"`
	ParkingSituation    bool            `json:"parking_situation" gorm:"column:parking_situation;"`
	wifi                bool            `json:"room_wifi" gorm:"column:room_wifi;"`
  CurfewSituation     bool            `json:"curfew_situation" gorm:"column:curfew_situation;"`
  ShareRoomAsLandlord bool            `json:"share_room_as_landlord" gorm:"column:share_room_as_landlord;"`
  RoomRefrigerator    bool            `json:"room_refrigerator;" gorm:"column:room_refrigerator;"`
  RoomWashingMachine  bool            `json:"room_washing_machine;" gorm:"column:room_washing_machine;"`
  securityGuard       bool            `json:"security_guard;" gorm:"column:security_guard;"`
  RoomBed             bool            `json:"room_bed;" gorm:"column:room_bed;"`
  RoomTivi            bool            `json:"room_tivi;" gorm:"column:room_tivi;"`
  RoomPetsAllowed     bool            `json:"room_pets_allowed;" gorm:"column:room_pets_allowed;"`
  RoomCloset          bool            `json:"room_closet;" gorm:"column:room_closet;"`
  RoomKitchen         bool            `json:"room_kitchen;" gorm:"column:room_kitchen;"`
  Window              bool            `json:"window;" gorm:"column:window;"`
  WaterHeater         bool            `json:"water_heater;" gorm:"column:water_heater;"`
  Loft                bool            `json:"loft;" gorm:"column:loft;"`
  Balcony             bool            `json:"balcony;" gorm:"column:balcony;"`
  ElectricPrice       int             `json:"electric_price" gorm:"electric_price;"`
  WaterPrice          int             `json:"water_price" gorm:"water_price;"`
  WifiCost            int             `json:"wifi_cost" gorm:"wifi_cost;"`
  RoomImages          *common.Images  `json:"room_images" gorm:"column:room_images;"`
  Status              int             `json:"status" gorm:"status;"`
  likedCount        int               `json:"like_count" gorm:"-"`
}
type HouseUpdate struct {
  common.SQLModel `json:",inline"`
  OwnerId             int             `json:"owner_id" gorm:"column:owner_id;"`
  Name                string          `json:"name" gorm:"column:name;"`
  Description         string          `json:"description" gorm:"description;"`
  Price               float64         `json:"price" gorm:"price;"`
  lat                 float64         `json:"lat" gorm:"column:lat;"`
  lng                 float64         `json:"lng" gorm:"column:lng;"`
  PhoneNumber         string          `json:"phone_number" gorm:"column:phone_number;"`
  airConditioner      bool            `json:"air_conditioner" gorm:"column:air_conditioner;"`
  RoomBathroom        bool            `json:"room_bathroom" gorm:"column:room_bathroom;"`
  ParkingSituation    bool            `json:"parking_situation" gorm:"column:parking_situation;"`
  wifi                bool            `json:"room_wifi" gorm:"column:room_wifi;"`
  CurfewSituation     bool            `json:"curfew_situation" gorm:"column:curfew_situation;"`
  ShareRoomAsLandlord bool            `json:"share_room_as_landlord" gorm:"column:share_room_as_landlord;"`
  RoomRefrigerator    bool            `json:"room_refrigerator;" gorm:"column:room_refrigerator;"`
  RoomWashingMachine  bool            `json:"room_washing_machine;" gorm:"column:room_washing_machine;"`
  securityGuard       bool            `json:"security_guard;" gorm:"column:security_guard;"`
  RoomBed             bool            `json:"room_bed;" gorm:"column:room_bed;"`
  RoomTivi            bool            `json:"room_tivi;" gorm:"column:room_tivi;"`
  RoomPetsAllowed     bool            `json:"room_pets_allowed;" gorm:"column:room_pets_allowed;"`
  RoomCloset          bool            `json:"room_closet;" gorm:"column:room_closet;"`
  RoomKitchen         bool            `json:"room_kitchen;" gorm:"column:room_kitchen;"`
  Window              bool            `json:"window;" gorm:"column:window;"`
  WaterHeater         bool            `json:"water_heater;" gorm:"column:water_heater;"`
  Loft                bool            `json:"loft;" gorm:"column:loft;"`
  Balcony             bool            `json:"balcony;" gorm:"column:balcony;"`
  ElectricPrice       int             `json:"electric_price" gorm:"electric_price;"`
  WaterPrice          int             `json:"water_price" gorm:"water_price;"`
  WifiCost            int             `json:"wifi_cost" gorm:"wifi_cost;"`
  RoomImages          *common.Images      `json:"room_images" gorm:"column:room_images;"`
  Status              int             `json:"status" gorm:"status;"`
}
type HouseCreate struct {
  common.SQLModel `json:",inline"`
  OwnerId             int             `json:"owner_id" gorm:"column:owner_id;"`
  Name                string          `json:"name" gorm:"column:name;"`
  Description         string          `json:"description" gorm:"description;"`
  Price               float64         `json:"price" gorm:"price;"`
  lat                 float64         `json:"lat" gorm:"column:lat;"`
  lng                 float64         `json:"lng" gorm:"column:lng;"`
  PhoneNumber         string          `json:"phone_number" gorm:"column:phone_number;"`
  airConditioner      bool            `json:"air_conditioner" gorm:"column:air_conditioner;"`
  RoomBathroom        bool            `json:"room_bathroom" gorm:"column:room_bathroom;"`
  ParkingSituation    bool            `json:"parking_situation" gorm:"column:parking_situation;"`
  wifi                bool            `json:"room_wifi" gorm:"column:room_wifi;"`
  CurfewSituation     bool            `json:"curfew_situation" gorm:"column:curfew_situation;"`
  ShareRoomAsLandlord bool            `json:"share_room_as_landlord" gorm:"column:share_room_as_landlord;"`
  RoomRefrigerator    bool            `json:"room_refrigerator;" gorm:"column:room_refrigerator;"`
  RoomWashingMachine  bool            `json:"room_washing_machine;" gorm:"column:room_washing_machine;"`
  securityGuard       bool            `json:"security_guard;" gorm:"column:security_guard;"`
  RoomBed             bool            `json:"room_bed;" gorm:"column:room_bed;"`
  RoomTivi            bool            `json:"room_tivi;" gorm:"column:room_tivi;"`
  RoomPetsAllowed     bool            `json:"room_pets_allowed;" gorm:"column:room_pets_allowed;"`
  RoomCloset          bool            `json:"room_closet;" gorm:"column:room_closet;"`
  RoomKitchen         bool            `json:"room_kitchen;" gorm:"column:room_kitchen;"`
  Window              bool            `json:"window;" gorm:"column:window;"`
  WaterHeater         bool            `json:"water_heater;" gorm:"column:water_heater;"`
  Loft                bool            `json:"loft;" gorm:"column:loft;"`
  Balcony             bool            `json:"balcony;" gorm:"column:balcony;"`
  ElectricPrice       int             `json:"electric_price" gorm:"electric_price;"`
  WaterPrice          int             `json:"water_price" gorm:"water_price;"`
  WifiCost            int             `json:"wifi_cost" gorm:"wifi_cost;"`
  RoomImages          *common.Images      `json:"room_images" gorm:"column:room_images;"`
  Status              int             `json:"status" gorm:"status;"`
}

func (HouseUpdate) TableName() string {
	return House{}.TableName()
}
func (House) TableName() string {
  return "houses"
}
func (HouseCreate) TableName() string {
	return House{}.TableName()
}
func (house *HouseCreate) Validate() error {
  house.Name = strings.TrimSpace(house.Name)
  if len(house.Name) == 0 {
    return ErrNameCannotBeEmpty
  }
  return nil
}
func (house *HouseUpdate) Validate() error {
  house.Name = strings.TrimSpace(house.Name)
  if len(house.Name) == 0 {
    return ErrNameCannotBeEmpty
  }
  return nil
}
var (
  ErrNameCannotBeEmpty = common.NewCustomError(nil, "house's name can't be blank", "ErrNameCannotBeEmpty")
)
func (product *House) Mask(isAdminOrOwner bool)  {
  product.GenerateUID(common.DbTypeRestaurant)
}
