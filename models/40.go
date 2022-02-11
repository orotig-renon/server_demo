package models

import "github.com/orotig/server_demo/config"

func GetWelderOrders(welder_orders *[]WelderTest) (err error) {
	if err = config.PGDB.Find(welder_orders).Error; err != nil {
		return err
	}
	return nil
}

func UpdateWelderOrder(welder_order *WelderTest) (err error) {
	if err = config.PGDB.Save(welder_order).Error; err != nil {
		return err
	}
	return nil
}
