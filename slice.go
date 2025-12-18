package main

// import "fmt"

// func main() {
// 	names := [...]string{"Laksmana", "Yudha", "Dede", "Budi", "Eko", "Kurniawan"}

// 	sliceSatu := names[4:6]
// 	fmt.Println(sliceSatu)

// 	sliceDua := names[:3]
// 	fmt.Println(sliceDua)

// 	sliceTiga := names[3:]
// 	fmt.Println(sliceTiga)

// 	var sliceEmpat []string = names[:]
// 	fmt.Println(sliceEmpat)

// 	var sliceLima []string = names[2:4]

// 	fmt.Println(len(sliceLima))
// 	fmt.Println(cap(sliceLima))

// 	sliceLima[0] = "Edited"
// 	fmt.Println(sliceLima)
// 	fmt.Println(names)

// 	sliceEnam := append(sliceLima, "Test")
// 	fmt.Println(sliceLima)
// 	fmt.Println(sliceEnam)
// 	fmt.Println(cap(sliceLima))
// 	fmt.Println(cap(sliceEnam))

// 	sliceEnam1 := append(sliceEnam, "Test 1")
// 	sliceEnam2 := append(sliceEnam, "Test 2")
// 	// sliceEnam3 := append(sliceEnam, "Test 3")
// 	// sliceEnam4 := append(sliceEnam, "Test 4")
// 	// sliceEnam5 := append(sliceEnam, "Test 5")

// 	fmt.Println(sliceEnam1)
// 	fmt.Println(sliceEnam2)
// 	// fmt.Println(sliceEnam3)
// 	// fmt.Println(sliceEnam4)
// 	// fmt.Println(sliceEnam5)
// 	fmt.Println(sliceEnam)

// 	sliceTujuh := append(sliceSatu, "Nilai Baru")
// 	fmt.Println(sliceSatu)
// 	fmt.Println(sliceTujuh)
// 	fmt.Println(names)

// 	sliceTujuh[0] = "Nilai Edited"
// 	fmt.Println(sliceSatu)
// 	fmt.Println(sliceTujuh)
// 	fmt.Println(names)

// 	// Key note: jika append sudah melebihi kapasitas maka slice akan dibuat baru, tidak mereferensi ke array yang lama
// 	// Key note: slice hanya mereferensi jika nilai pada slice diubah / append (selama tidak melebihi kapasitas), array asli juga akan ikut berubah
// 	// Key note: len => panjang slice, capacity => kapasitas slice yaitu index start slice hingga index akhir array
// 	// Keyn note: append mereferensi setelah index terakhir slice (pada array bisa index ke n), jika pada array masih ada next elemen, maka bisa jadi akan direplace

// 	// ini akan membuat array dibalik layar dengan kapasitas 5, lalu langsung buat slice sepanjang 2 mulai dari index 0
// 	var newSlice []string = make([]string, 2, 5)
// 	newSlice[0] = "Yudha"
// 	newSlice[1] = "Yudha"
// 	// newSlice[2] = "Yudha"

// 	fmt.Println(newSlice)

// 	newSlice2 := append(newSlice, "YudhaL")
// 	fmt.Println(newSlice)
// 	fmt.Println(newSlice2)

// 	newSlice2[0] = "Dede"
// 	fmt.Println(newSlice)
// 	fmt.Println(newSlice2)

// 	// array hari
// 	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabut", "Minggu"}
// 	fromSlice := days[:]

// 	// array lain yang masih kosong
// 	toSlice := make([]string, len(fromSlice), cap(fromSlice))

// 	copy(toSlice, fromSlice)
// 	fromSlice[0] = "Hari Libur" // tidak ngefek karena beda array

// 	fmt.Println(fromSlice)
// 	fmt.Println(toSlice)
// }