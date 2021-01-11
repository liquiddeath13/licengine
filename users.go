package main

import (
	"io/ioutil"
	"time"
)

func (s *subscription) isZero() bool {
	return s.Key == ""
}

func (s *subscription) daysTillEnd() int {
	return s.SubDays - int(s.Created.Sub(time.Now()).Hours()/24)
}

func findSub(key string) subscription {
	for _, u := range users {
		u.DaysTillEnd = u.daysTillEnd()
		if u.Key == key && u.DaysTillEnd > 0 {
			return u
		}
	}
	return subscription{}
}

func setLastLogin(s subscription) {
	for _, u := range users {
		if u.Key == s.Key {
			u.LastLoginDate = time.Now()
			break
		}
	}
}

func getLibByName(libName string) []byte {
	for _, sd := range warez {
		if sd.Name == libName {
			content, _ := ioutil.ReadFile(sd.Internal)
			return content
		}
	}
	return nil
}

func getAvailablesWareList(s subscription) []softwareDescription {
	availableWarezList := []softwareDescription{}
	for _, sd := range warez {
		if isContains(s.AvailableWarezIDs, sd.ID) {
			availableWarezList = append(availableWarezList, sd)
		}
	}
	return availableWarezList
}
