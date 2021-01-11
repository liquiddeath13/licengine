package main

import "time"

func encryptDecrypt(input, key string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}
	return output
}

func isContains(list []int, e int) bool {
	for _, v := range list {
		if e == v {
			return true
		}
	}
	return false
}

func addUser(name string, key string, subDays int, LastLoginDate time.Time, Created time.Time, availableWarezIDs []int) {
	s := subscription{}
	s.Name = name
	s.Key = key
	s.SubDays = subDays
	s.LastLoginDate = LastLoginDate
	s.Created = Created
	s.AvailableWarezIDs = availableWarezIDs
	users = append(users, s)
}

func addWare(ID int, name string, internal string) {
	sd := softwareDescription{}
	sd.ID = ID
	sd.Name = name
	sd.Internal = internal
	warez = append(warez, sd)
}

func initUsers() {
	addUser("admin", "admin", 365, time.Now(), time.Now(), []int{1, 2, 3})
	addUser("test1", "1test", 365, time.Now(), time.Now(), []int{1})
	addUser("test2", "2test", 365, time.Now(), time.Now(), []int{2})
	addUser("test3", "3test", 365, time.Now(), time.Now(), []int{3})
	addUser("test12", "12test", 365, time.Now(), time.Now(), []int{1, 2})
	addUser("test13", "13test", 365, time.Now(), time.Now(), []int{1, 3})
	addUser("test23", "23test", 365, time.Now(), time.Now(), []int{2, 3})
}

func initWareList() {
	addWare(1, "Some Software 1", "somesoftware1")
	addWare(2, "Some Software 2", "somesoftware2")
	addWare(3, "Some Software 3", "somesoftware3")
}
