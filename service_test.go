package keycloakmiddleware

import (
	"testing"
)

func Benchmark_isScopesValid_300scope1000claims(b *testing.B) {
	c, scopes := get300scope1000claims()
	for i := 0; i < b.N; i++ {
		isScopesValid(c, scopes)
	}
}

func Benchmark_isScopesValidNew_300scope1000claims(b *testing.B) {
	c, scopes := get300scope1000claims()
	for i := 0; i < b.N; i++ {
		isScopesValidNew(c, scopes)
	}
}

func Benchmark_isScopesValid_1000scope300claims(b *testing.B) {
	c, scopes := get1000scope300claims()
	for i := 0; i < b.N; i++ {
		isScopesValid(c, scopes)
	}
}

func Benchmark_isScopesValidNew_1000scope300claims(b *testing.B) {
	c, scopes := get1000scope300claims()
	for i := 0; i < b.N; i++ {
		isScopesValidNew(c, scopes)
	}
}

func get300scope1000claims() (claims, []string) {
	claims := claims{
		Authorization: authorization{
			Permissions: []permission{
				{Scopes: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "90", "91", "92", "93", "94", "95", "96", "97", "98", "99", "100"}},
				{Scopes: []string{"101", "102", "103", "104", "105", "106", "107", "108", "109", "110", "111", "112", "113", "114", "115", "116", "117", "118", "119", "120", "121", "122", "123", "124", "125", "126", "127", "128", "129", "130", "131", "132", "133", "134", "135", "136", "137", "138", "139", "140", "141", "142", "143", "144", "145", "146", "147", "148", "149", "150", "151", "152", "153", "154", "155", "156", "157", "158", "159", "160", "161", "162", "163", "164", "165", "166", "167", "168", "169", "170", "171", "172", "173", "174", "175", "176", "177", "178", "179", "180", "181", "182", "183", "184", "185", "186", "187", "188", "189", "190", "191", "192", "193", "194", "195", "196", "197", "198", "199", "200"}},
				{Scopes: []string{"201", "202", "203", "204", "205", "206", "207", "208", "209", "210", "211", "212", "213", "214", "215", "216", "217", "218", "219", "220", "221", "222", "223", "224", "225", "226", "227", "228", "229", "230", "231", "232", "233", "234", "235", "236", "237", "238", "239", "240", "241", "242", "243", "244", "245", "246", "247", "248", "249", "250", "251", "252", "253", "254", "255", "256", "257", "258", "259", "260", "261", "262", "263", "264", "265", "266", "267", "268", "269", "270", "271", "272", "273", "274", "275", "276", "277", "278", "279", "280", "281", "282", "283", "284", "285", "286", "287", "288", "289", "290", "291", "292", "293", "294", "295", "296", "297", "298", "299", "300"}},
				{Scopes: []string{"301", "302", "303", "304", "305", "306", "307", "308", "309", "310", "311", "312", "313", "314", "315", "316", "317", "318", "319", "320", "321", "322", "323", "324", "325", "326", "327", "328", "329", "330", "331", "332", "333", "334", "335", "336", "337", "338", "339", "340", "341", "342", "343", "344", "345", "346", "347", "348", "349", "350", "351", "352", "353", "354", "355", "356", "357", "358", "359", "360", "361", "362", "363", "364", "365", "366", "367", "368", "369", "370", "371", "372", "373", "374", "375", "376", "377", "378", "379", "380", "381", "382", "383", "384", "385", "386", "387", "388", "389", "390", "391", "392", "393", "394", "395", "396", "397", "398", "399", "400"}},
				{Scopes: []string{"401", "402", "403", "404", "405", "406", "407", "408", "409", "410", "411", "412", "413", "414", "415", "416", "417", "418", "419", "420", "421", "422", "423", "424", "425", "426", "427", "428", "429", "430", "431", "432", "433", "434", "435", "436", "437", "438", "439", "440", "441", "442", "443", "444", "445", "446", "447", "448", "449", "450", "451", "452", "453", "454", "455", "456", "457", "458", "459", "460", "461", "462", "463", "464", "465", "466", "467", "468", "469", "470", "471", "472", "473", "474", "475", "476", "477", "478", "479", "480", "481", "482", "483", "484", "485", "486", "487", "488", "489", "490", "491", "492", "493", "494", "495", "496", "497", "498", "499", "500"}},
				{Scopes: []string{"501", "502", "503", "504", "505", "506", "507", "508", "509", "510", "511", "512", "513", "514", "515", "516", "517", "518", "519", "520", "521", "522", "523", "524", "525", "526", "527", "528", "529", "530", "531", "532", "533", "534", "535", "536", "537", "538", "539", "540", "541", "542", "543", "544", "545", "546", "547", "548", "549", "550", "551", "552", "553", "554", "555", "556", "557", "558", "559", "560", "561", "562", "563", "564", "565", "566", "567", "568", "569", "570", "571", "572", "573", "574", "575", "576", "577", "578", "579", "580", "581", "582", "583", "584", "585", "586", "587", "588", "589", "590", "591", "592", "593", "594", "595", "596", "597", "598", "599", "600"}},
				{Scopes: []string{"601", "602", "603", "604", "605", "606", "607", "608", "609", "610", "611", "612", "613", "614", "615", "616", "617", "618", "619", "620", "621", "622", "623", "624", "625", "626", "627", "628", "629", "630", "631", "632", "633", "634", "635", "636", "637", "638", "639", "640", "641", "642", "643", "644", "645", "646", "647", "648", "649", "650", "651", "652", "653", "654", "655", "656", "657", "658", "659", "660", "661", "662", "663", "664", "665", "666", "667", "668", "669", "670", "671", "672", "673", "674", "675", "676", "677", "678", "679", "680", "681", "682", "683", "684", "685", "686", "687", "688", "689", "690", "691", "692", "693", "694", "695", "696", "697", "698", "699", "700"}},
				{Scopes: []string{"701", "702", "703", "704", "705", "706", "707", "708", "709", "710", "711", "712", "713", "714", "715", "716", "717", "718", "719", "720", "721", "722", "723", "724", "725", "726", "727", "728", "729", "730", "731", "732", "733", "734", "735", "736", "737", "738", "739", "740", "741", "742", "743", "744", "745", "746", "747", "748", "749", "750", "751", "752", "753", "754", "755", "756", "757", "758", "759", "760", "761", "762", "763", "764", "765", "766", "767", "768", "769", "770", "771", "772", "773", "774", "775", "776", "777", "778", "779", "780", "781", "782", "783", "784", "785", "786", "787", "788", "789", "790", "791", "792", "793", "794", "795", "796", "797", "798", "799", "800"}},
				{Scopes: []string{"801", "802", "803", "804", "805", "806", "807", "808", "809", "810", "811", "812", "813", "814", "815", "816", "817", "818", "819", "820", "821", "822", "823", "824", "825", "826", "827", "828", "829", "830", "831", "832", "833", "834", "835", "836", "837", "838", "839", "840", "841", "842", "843", "844", "845", "846", "847", "848", "849", "850", "851", "852", "853", "854", "855", "856", "857", "858", "859", "860", "861", "862", "863", "864", "865", "866", "867", "868", "869", "870", "871", "872", "873", "874", "875", "876", "877", "878", "879", "880", "881", "882", "883", "884", "885", "886", "887", "888", "889", "890", "891", "892", "893", "894", "895", "896", "897", "898", "899", "900"}},
				{Scopes: []string{"901", "902", "903", "904", "905", "906", "907", "908", "909", "910", "911", "912", "913", "914", "915", "916", "917", "918", "919", "920", "921", "922", "923", "924", "925", "926", "927", "928", "929", "930", "931", "932", "933", "934", "935", "936", "937", "938", "939", "940", "941", "942", "943", "944", "945", "946", "947", "948", "949", "950", "951", "952", "953", "954", "955", "956", "957", "958", "959", "960", "961", "962", "963", "964", "965", "966", "967", "968", "969", "970", "971", "972", "973", "974", "975", "976", "977", "978", "979", "980", "981", "982", "983", "984", "985", "986", "987", "988", "989", "990", "991", "992", "993", "994", "995", "996", "997", "998", "999", "1000"}},
			},
		},
	}

	scopes := []string{
		"1-ajMOt", "2-loEWh", "3-cuTsu", "4-lSiRP", "5-pwdkO", "6-EdDAE", "7-ngjhQ", "8-cCwFt", "9-smiAM", "10-snjdB", "11-CMHWW", "12-TNIxI", "13-CeVyu", "14-jJDdE", "15-Vhvaa", "16-kYOAU", "17-NlPKN", "18-QOQxX", "19-aLNHC", "20-NZHAx", "21-HYAyu", "22-clMAZ", "23-ofgCd", "24-tFZYP", "25-ejQKx", "26-lfPgC", "27-cSNtp", "28-xzihw", "29-iUvpG", "30-WSgaX", "31-TzJKv", "32-BlIoT", "33-IWPfu", "34-UuGeR", "35-XCcSE", "36-Xbpmc", "37-lYPzu", "38-eeVFL", "39-exJqU", "40-MyyPQ", "41-NQXTo", "42-RGrph", "43-pHoTC", "44-MtCmq", "45-hzHSX", "46-GTkLq", "47-VDrjF", "48-KUcIP", "49-VSFaM", "50-rYBbB", "51-IkRaG", "52-CcoKv", "53-RwfES", "54-dllwm", "55-JbVIX", "56-TTUsz", "57-CgZKU", "58-TGbpn", "59-jPnOi", "60-UcVNf", "61-agDbz", "62-tQzQQ", "63-ERHVo", "64-lChyX", "65-LRVMB", "66-SNpaB", "67-RBRjt", "68-GAzDn", "69-qYCij", "70-Ozlle", "71-KzFTy", "72-ouYGD", "73-HcMmP", "74-KtKln", "75-usYJi", "76-CPcaF", "77-BOOte", "78-XgzNS", "79-HedJC", "80-DidHV", "81-vUeQU", "82-MOYbF", "83-ZQITg", "84-pdNNi", "85-XjTDG", "86-RfkFc", "87-DTIKO", "88-UIKbY", "89-MkptK", "90-ObHaj", "91-euawt", "92-fUyvh", "93-QOEUv", "94-Afqbl", "95-ysGNZ", "96-xYSXs", "97-cANHL", "98-Hwiti", "99-ljIXr", "100-jPoxC",
		"101-TtHvl", "102-OodsB", "103-YrItq", "104-tXNAi", "105-TVUUI", "106-OpCwL", "107-UcDYG", "108-rEvTG", "109-HCzqe", "110-kepwW", "111-GYnEI", "112-WsIAX", "113-UTMTc", "114-mJYnN", "115-gJeWg", "116-kTttx", "117-LCXyx", "118-ckwRj", "119-yVvBz", "120-otIRS", "121-RxEtY", "122-DuZoY", "123-MPlkC", "124-zDCnt", "125-DJreE", "126-Gwecc", "127-dqqmB", "128-nFrfM", "129-wlTTP", "130-mguDp", "131-HVNeT", "132-uIsnO", "133-ShmZn", "134-VCGDr", "135-vKmfZ", "136-ddOph", "137-lMqSg", "138-MjYXX", "139-APkKW", "140-qWekp", "141-fbkWr", "142-FHyrx", "143-AsyCI", "144-djOcn", "145-OqICO", "146-ESbAu", "147-gKtAK", "148-XeKfY", "149-GwSVj", "150-KxKED", "151-aqWFj", "152-ULwkR", "153-zKgLl", "154-IbIyX", "155-HVJtu", "156-lMzfj", "157-eIyPR", "158-UFmdD", "159-Xzlar", "160-zHnnt", "161-iOCCE", "162-YnbRt", "163-vcCEk", "164-cEUEy", "165-Zkzcq", "166-KndyK", "167-Ymfpf", "168-znqzw", "169-oLiDl", "170-vJgnj", "171-nSmIT", "172-Rxymy", "173-GbeTX", "174-yseBm", "175-osXtD", "176-OaMLX", "177-zZgWA", "178-AlfIa", "179-icjqo", "180-rHtGc", "181-sZzOg", "182-JrUHw", "183-CUglz", "184-nVHQv", "185-BqVqt", "186-FGaZY", "187-UJmAH", "188-EJMxv", "189-fXzPN", "190-eHQOR", "191-TaoIt", "192-BwqXc", "193-RypNx", "194-RgrHI", "195-rossF", "196-ygAKL", "197-komEq", "198-jfdvu", "199-xLCox", "200-PynwV",
		"201-aMUim", "202-oSfpA", "203-YFMOy", "204-mUerQ", "205-SSHXo", "206-kPYEQ", "207-xozUZ", "208-XHlQy", "209-pBNYD", "210-hJUdL", "211-tdeeU", "212-tTWHg", "213-BvcqT", "214-YfHfl", "215-kuHij", "216-ETGwO", "217-TbaIM", "218-nlYwR", "219-ASaiH", "220-DpJfk", "221-MvRNN", "222-ZxfME", "223-yyPvw", "224-xkkmC", "225-BVVQh", "226-dowuz", "227-ShLoV", "228-xyaQJ", "229-VQCIr", "230-KIQkH", "231-MsJOI", "232-AGEVG", "233-xOblX", "234-xGkKE", "235-KVrQx", "236-bEwsk", "237-XgHrA", "238-BdCFJ", "239-RaBdD", "240-xEoXv", "241-HHNMH", "242-wzvyd", "243-awCQj", "244-YiNzg", "245-bdBGG", "246-lXpEp", "247-pZBvI", "248-dnjDS", "249-FAHhg", "250-CXRMB", "251-YUqZJ", "252-eGbos", "253-FplYL", "254-UBxQK", "255-HZJIJ", "256-kbivn", "257-MPVAX", "258-NCPsc", "259-vNhTH", "260-LfmtW", "261-HkgDn", "262-CXeup", "263-lKCGM", "264-MyYMM", "265-EXrmE", "266-FEzBb", "267-dDvhj", "268-HmIeQ", "269-zKZoB", "270-VSZEU", "271-zxAZm", "272-AMMEP", "273-pMxuv", "274-UMvTg", "275-YaWQo", "276-TMNpB", "277-bpphM", "278-eAvoM", "279-IuCdq", "280-zwFxL", "281-KwfBv", "282-TWYMf", "283-DvVuj", "284-RNRzX", "285-sKPAM", "286-bDTVr", "287-AeGom", "288-slhBb", "289-jokIh", "290-rFBnw", "291-FfLgw", "292-mUEqC", "293-aeRKE", "294-zwtve", "295-IcZRl", "296-pDVRR", "297-gAexw", "298-AqsRD", "299-Qqilx", "300-yHLRp",
	}

	return claims, scopes
}

func get1000scope300claims() (claims, []string) {
	claims := claims{
		Authorization: authorization{
			Permissions: []permission{
				{Scopes: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70", "71", "72", "73", "74", "75", "76", "77", "78", "79", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89", "90", "91", "92", "93", "94", "95", "96", "97", "98", "99", "100"}},
				{Scopes: []string{"101", "102", "103", "104", "105", "106", "107", "108", "109", "110", "111", "112", "113", "114", "115", "116", "117", "118", "119", "120", "121", "122", "123", "124", "125", "126", "127", "128", "129", "130", "131", "132", "133", "134", "135", "136", "137", "138", "139", "140", "141", "142", "143", "144", "145", "146", "147", "148", "149", "150", "151", "152", "153", "154", "155", "156", "157", "158", "159", "160", "161", "162", "163", "164", "165", "166", "167", "168", "169", "170", "171", "172", "173", "174", "175", "176", "177", "178", "179", "180", "181", "182", "183", "184", "185", "186", "187", "188", "189", "190", "191", "192", "193", "194", "195", "196", "197", "198", "199", "200"}},
				{Scopes: []string{"201", "202", "203", "204", "205", "206", "207", "208", "209", "210", "211", "212", "213", "214", "215", "216", "217", "218", "219", "220", "221", "222", "223", "224", "225", "226", "227", "228", "229", "230", "231", "232", "233", "234", "235", "236", "237", "238", "239", "240", "241", "242", "243", "244", "245", "246", "247", "248", "249", "250", "251", "252", "253", "254", "255", "256", "257", "258", "259", "260", "261", "262", "263", "264", "265", "266", "267", "268", "269", "270", "271", "272", "273", "274", "275", "276", "277", "278", "279", "280", "281", "282", "283", "284", "285", "286", "287", "288", "289", "290", "291", "292", "293", "294", "295", "296", "297", "298", "299", "300"}},
			},
		},
	}

	scopes := []string{
		"1-tKdKU", "2-UGmfL", "3-AtlSx", "4-LhiqL", "5-myJlO", "6-oJjlm", "7-xHrAi", "8-muYHw", "9-eYOJo", "10-TQGTu", "11-SikMS", "12-xkKqi", "13-NmTkZ", "14-pFIIb", "15-hmDxT", "16-EwSgo", "17-fagUC", "18-QpQTf", "19-vyCUJ", "20-QmdSz", "21-wSyjG", "22-UFHvG", "23-FKeKa", "24-yCEjW", "25-Ecivw", "26-aDhDQ", "27-KYIQd", "28-THmJt", "29-YcTkr", "30-LASdT", "31-ZWXDK", "32-DnLAr", "33-zmStF", "34-ilUgI", "35-YcVnL", "36-rrCEL", "37-BJYKW", "38-rfKrA", "39-OkqiC", "40-tODLf", "41-gtiFG", "42-fLiEb", "43-FPWZX", "44-RqAMO", "45-zAamo", "46-gqTNF", "47-IcXkO", "48-VQXMr", "49-ECSbY", "50-QtjLg", "51-gaflS", "52-CPesh", "53-XLwEr", "54-YcYpt", "55-fVWDT", "56-oGjeR", "57-nPzGe", "58-SqoLJ", "59-QSsSb", "60-OwrDO", "61-CTBdj", "62-WKHlY", "63-pwWxP", "64-EYRxQ", "65-vhDHH", "66-Nzctt", "67-DsKxx", "68-ludWG", "69-OQZJS", "70-EgJOs", "71-zcYPF", "72-MCXjL", "73-QwkQf", "74-cUFtr", "75-jIllN", "76-lxReY", "77-hnyQV", "78-bbQoO", "79-ZrBXn", "80-CLffB", "81-IMSgQ", "82-xavDi", "83-FKnTm", "84-ZLith", "85-LWwci", "86-xpyWE", "87-tELGi", "88-YHtRH", "89-dISII", "90-gduuN", "91-APOIH", "92-SePtq", "93-GdcaE", "94-ceuOI", "95-lffpc", "96-Xqgxq", "97-MFvaB", "98-cqiVv", "99-nHmpv", "100-WQTMt",
		"101-EiVWH", "102-wAoel", "103-ZdBtB", "104-otNxA", "105-dndKt", "106-codiK", "107-VmVLx", "108-noIHO", "109-jBgOo", "110-LeSuG", "111-sgWGx", "112-MbHiV", "113-pJwor", "114-QAxcH", "115-gfTYR", "116-rYwHO", "117-tpNvp", "118-yWsLh", "119-eswRw", "120-pqqrg", "121-UFpjG", "122-vdTwh", "123-gGXyy", "124-NXLTP", "125-MzATx", "126-BGTux", "127-DMvmd", "128-ECpPn", "129-mFKPZ", "130-ezFHz", "131-AtCug", "132-PKajx", "133-knNOY", "134-VZuAl", "135-nYTZF", "136-NoMSb", "137-ExDtu", "138-NQciT", "139-LanaM", "140-psaDM", "141-QOrfP", "142-qknAa", "143-DgGTj", "144-ZrJlv", "145-srEcl", "146-brbGo", "147-AcOKH", "148-ONmAA", "149-oWOxu", "150-ROTcB", "151-dAKKw", "152-GuEtp", "153-LqToA", "154-WejVS", "155-XTsuR", "156-qooXz", "157-qAXMu", "158-DtFOs", "159-SojWq", "160-iyLji", "161-QVeTW", "162-bTWUB", "163-DIeow", "164-acfHK", "165-xKoFq", "166-cJjbK", "167-dNiCc", "168-luOYy", "169-aHRmt", "170-ZnHLK", "171-Ldbkr", "172-qdqNY", "173-mixTL", "174-KghOf", "175-cDxDN", "176-qCGcl", "177-sToXL", "178-YwXlj", "179-DXwHd", "180-XxaxM", "181-qVbzB", "182-tagYd", "183-TkDOc", "184-USLOM", "185-tUyFe", "186-WKKIM", "187-QOvnG", "188-QmBJj", "189-ScJFq", "190-wWMce", "191-iCMVQ", "192-IddxU", "193-JMnlD", "194-nANzF", "195-RtyFU", "196-aoMof", "197-bjpJB", "198-NmdcD", "199-EHdlg", "200-BZOfe",
		"201-fonyQ", "202-xNBWH", "203-kbngy", "204-fmPzY", "205-RPsDY", "206-QTfjK", "207-AJqZW", "208-FtJum", "209-lpjqc", "210-FNNNU", "211-OOZRE", "212-HwNok", "213-IicYI", "214-wpsgs", "215-RACsJ", "216-LisIU", "217-fcEte", "218-MEwAv", "219-oOOlF", "220-gyfNe", "221-cvZrD", "222-BFtwC", "223-dxVsu", "224-rqLyw", "225-ynxQd", "226-lXEWe", "227-AkJzD", "228-lerga", "229-vLkpd", "230-vGLAs", "231-VGHDx", "232-AKAUq", "233-vaNsE", "234-LxkvE", "235-UVImF", "236-GFjhc", "237-ZmJmW", "238-boZJT", "239-nDKbX", "240-RtyGF", "241-mHCZH", "242-yaBOy", "243-nlkAy", "244-eIMJW", "245-WmJwM", "246-XqeNX", "247-GBxty", "248-OYWzB", "249-JDGuz", "250-rDtpZ", "251-pvpex", "252-bAYZU", "253-wFPhB", "254-ezeoy", "255-BVVuw", "256-JZmeI", "257-uJbZD", "258-HUJEe", "259-YNPUZ", "260-RZVls", "261-gSCVy", "262-PQXma", "263-nruoH", "264-qJPov", "265-KzSsU", "266-VIqOQ", "267-FnOZI", "268-StZvY", "269-SzChw", "270-gCBSb", "271-JfVId", "272-CZHih", "273-mTTJf", "274-SADvI", "275-XiptG", "276-YwtPz", "277-ArGFu", "278-EWXsQ", "279-OPNFZ", "280-YSYDB", "281-HEJVH", "282-ljtXO", "283-HgnCl", "284-fOLNW", "285-uGkVb", "286-UQJTb", "287-qhAhW", "288-KzYga", "289-Iillf", "290-wGeoi", "291-AGNgm", "292-JQajt", "293-GnMsf", "294-PknFK", "295-ccNbn", "296-WHMNQ", "297-ZNOCa", "298-XdXax", "299-KjWPe", "300-GSUkx",
		"301-AvQRO", "302-oGhoA", "303-rrBqU", "304-Wcuhs", "305-JSgZP", "306-cFOeN", "307-MTbSa", "308-jgZAK", "309-WPIYO", "310-imntX", "311-JEylz", "312-UNHgg", "313-mcmNN", "314-SVZdL", "315-PUeUS", "316-kxuRK", "317-cGdnP", "318-dbpGX", "319-eNUSr", "320-duefV", "321-iiaXv", "322-akfyL", "323-QuYoU", "324-kEbfc", "325-tlfiR", "326-aIBHF", "327-twSfz", "328-svYqI", "329-aTJWp", "330-TFmpK", "331-dZIUZ", "332-OmyZt", "333-bUkIu", "334-GNhzR", "335-EUxmG", "336-YCPWt", "337-YPEVE", "338-gkPcX", "339-qVykL", "340-koqWE", "341-pyJJg", "342-szENM", "343-hJSXA", "344-FPiVl", "345-gRABX", "346-DXLIl", "347-uLkzb", "348-thRqd", "349-Pdspz", "350-ajOrb", "351-HGoWA", "352-HSngs", "353-WOLlJ", "354-mdCdp", "355-UEcdT", "356-zIPko", "357-kQJkn", "358-ozbva", "359-bNfYc", "360-BItDB", "361-CEwZK", "362-IqiJo", "363-DGacZ", "364-tNWbN", "365-zqMap", "366-dULbU", "367-GQnGK", "368-cIGDc", "369-UnxMl", "370-sBlfY", "371-SUXnD", "372-jtWHg", "373-IqVzi", "374-jIxkt", "375-zSUbR", "376-fERAy", "377-sZFcO", "378-pALTp", "379-bDFHk", "380-KorCf", "381-jYPVL", "382-UXnOF", "383-Vjbhx", "384-tMuQa", "385-mBfmR", "386-TPSLX", "387-pwETX", "388-tPzgO", "389-iOHzT", "390-aISQO", "391-OJTmv", "392-CTFcW", "393-DsfrO", "394-SJOPD", "395-nvfUO", "396-uTqHA", "397-VxVpR", "398-mwiXT", "399-OGonb", "400-oGbQh",
		"401-LXJyY", "402-vVflG", "403-gUgvj", "404-oZeBt", "405-WsurX", "406-PeCJX", "407-IKfJL", "408-ZEKEl", "409-iSvoK", "410-zsrlH", "411-BwRBT", "412-VViZW", "413-qzUIm", "414-DPOSa", "415-OeSFM", "416-ZYsbt", "417-orpJB", "418-BusxT", "419-BcfHQ", "420-gSkAJ", "421-Swtnx", "422-wTBjD", "423-QswHu", "424-izdYD", "425-covhe", "426-ZIgIC", "427-hstwh", "428-dUuLs", "429-gHbcl", "430-LjmlG", "431-ZANft", "432-lpjzd", "433-PwDrZ", "434-iSFZm", "435-PFFFr", "436-QjlUp", "437-rgBOu", "438-IwZEg", "439-NCUXr", "440-gcvVc", "441-XsaVF", "442-TwjrB", "443-UEsOq", "444-ajVkI", "445-bnoer", "446-Opmot", "447-biVDL", "448-bUApo", "449-fSrmp", "450-vkngm", "451-WXyXY", "452-CezJY", "453-lKCkm", "454-UisJX", "455-PABOk", "456-PXWAX", "457-DzChF", "458-AeCcT", "459-Kyxkb", "460-vLnXh", "461-JsSPD", "462-syFPz", "463-FxChu", "464-DaKQn", "465-XcvfN", "466-SljdU", "467-IrjtM", "468-BqeZb", "469-bIipk", "470-bpKSc", "471-nnAnO", "472-zmKGa", "473-eGShw", "474-xyocn", "475-vDBvW", "476-uQkyW", "477-FjKrd", "478-cZMpv", "479-cLxrN", "480-jXTbk", "481-WYRdg", "482-tigcm", "483-DZUbn", "484-IuPdV", "485-TvwIC", "486-eHXTr", "487-OKVtP", "488-kcJTm", "489-tyaDJ", "490-JwSMc", "491-eXvsa", "492-mMnit", "493-DHDiO", "494-GaPyT", "495-txLSe", "496-hRBhB", "497-EyYxM", "498-ZERQA", "499-Hxpum", "500-uBFQw",
		"501-FhHUV", "502-tSJEG", "503-lkBjN", "504-KjBLD", "505-xVlXy", "506-oWmJx", "507-KAFTD", "508-mAvKX", "509-PKpbY", "510-IceOI", "511-zwXya", "512-JrCBE", "513-cjgvW", "514-YSKRK", "515-IhGgS", "516-gQLnx", "517-cJkNi", "518-ucMui", "519-ZlcYP", "520-FWqAj", "521-EBXvE", "522-siqAU", "523-qUpid", "524-gFvlD", "525-gzVPO", "526-QGMTQ", "527-mxOkd", "528-qrvNC", "529-zFFDr", "530-siOuR", "531-nidJj", "532-cMhQy", "533-BcHrf", "534-mKbgi", "535-YCUXI", "536-elEpk", "537-FLkqi", "538-lXHau", "539-NQROH", "540-ZJjYS", "541-TGSjc", "542-ZYEtz", "543-lhkiB", "544-UNhif", "545-vAjdc", "546-joOIG", "547-kzmTV", "548-Gonfl", "549-CLgkN", "550-vCJVe", "551-ctxOa", "552-GLRXb", "553-qOGAY", "554-LUiXz", "555-ycHmi", "556-Vmmxj", "557-tQjDy", "558-NFVen", "559-BaqIi", "560-FbOcP", "561-NNwEN", "562-RtQrT", "563-ClETX", "564-KGdfm", "565-CZiZi", "566-zOdKC", "567-pnFcJ", "568-tTxCl", "569-HIOSI", "570-YMobn", "571-msYKs", "572-xCUZS", "573-dDfOn", "574-jUhTV", "575-MMqjR", "576-iNguP", "577-uUdQh", "578-wyzdy", "579-JtshH", "580-vJDCS", "581-ZtRfi", "582-tqTzt", "583-Rvakn", "584-KgujG", "585-RRsvG", "586-MaarZ", "587-hqZfx", "588-sdBBn", "589-UlyQi", "590-vhrjR", "591-Dieaj", "592-yqTkd", "593-IJVHm", "594-CPSHG", "595-rnitg", "596-LmhRj", "597-oDQvW", "598-YtMgS", "599-lptfN", "600-DwxFD",
		"601-dgETs", "602-lFEoz", "603-vllnr", "604-gtZay", "605-UGqfT", "606-OJMhN", "607-LZOjR", "608-UELCC", "609-wWTlI", "610-tfaPZ", "611-djGsV", "612-LlYkx", "613-VZRrk", "614-ZaHJY", "615-DiXJp", "616-oCXbe", "617-rwmnP", "618-RsiUO", "619-xPVLD", "620-OJHvn", "621-fecPR", "622-RrULo", "623-amdja", "624-Zaxdn", "625-oJGjU", "626-SaUCV", "627-UvhVL", "628-rbKCK", "629-YVRUd", "630-bxvNS", "631-gQbIt", "632-fHFNp", "633-MkTcI", "634-yhFHb", "635-GxBDO", "636-rwMpt", "637-eNYqR", "638-nsPZH", "639-pZsaw", "640-fsQWW", "641-rotZq", "642-xBHGX", "643-OEita", "644-oPsHj", "645-NMuqH", "646-bugSm", "647-swwrR", "648-GGgAl", "649-VIEBk", "650-dLMTn", "651-XNuBE", "652-XmGYy", "653-SPYLu", "654-YKubM", "655-yAJpK", "656-dkFic", "657-WjfiS", "658-xyxpQ", "659-SQIVJ", "660-MYmDq", "661-DfMke", "662-TSZML", "663-HFIzu", "664-TNyhW", "665-Ecfoo", "666-sYJMY", "667-mARTq", "668-AMabQ", "669-weJZc", "670-cuOjk", "671-RzXQP", "672-ouzoW", "673-yCznZ", "674-UWmIJ", "675-qOjId", "676-wFBEC", "677-QkdGo", "678-RbkSs", "679-ePkNN", "680-AnrDl", "681-GnCWy", "682-OMQhz", "683-nxIUW", "684-YukVx", "685-Ijfjp", "686-PXnBR", "687-nfMfa", "688-pQURy", "689-paWPY", "690-TLUxW", "691-PIJxx", "692-AoUKo", "693-pNPaE", "694-IMjRV", "695-OHdbx", "696-FFvmz", "697-ynpno", "698-vknrW", "699-oBZXx", "700-enDDq",
		"701-YRPcr", "702-zepwc", "703-JaFqE", "704-wiJlb", "705-FmfVS", "706-WLyYJ", "707-nLWtv", "708-aZfGA", "709-INuGp", "710-YYgXc", "711-HYtCU", "712-CkYkb", "713-oiaEs", "714-FNvxb", "715-nDlib", "716-hzTQa", "717-YSbcr", "718-JwDQL", "719-DECKY", "720-HvSzQ", "721-lFTGg", "722-aBCyf", "723-eTqwu", "724-tNGpN", "725-utVgR", "726-ntGDJ", "727-mVXyN", "728-Ssfvt", "729-qBDxe", "730-bIWPI", "731-oVIQR", "732-hEIMp", "733-uHjor", "734-BQkEf", "735-ofjWY", "736-YoIMh", "737-FpqUV", "738-UFUDu", "739-nXVNj", "740-Yoycy", "741-PBABA", "742-iYrWx", "743-tvnor", "744-TxiXr", "745-IUaSB", "746-NXeOG", "747-BPpIM", "748-iKLDc", "749-LgDqH", "750-yftgW", "751-HeRVa", "752-IZnGB", "753-ZlggS", "754-Scfjb", "755-ISVkS", "756-FPpla", "757-bBejP", "758-sYJCX", "759-mpjtW", "760-JGkVq", "761-VkapZ", "762-bTNum", "763-Nvcgy", "764-Tbsnp", "765-FQyWA", "766-FAhVm", "767-vShDQ", "768-yMOyk", "769-jHaQP", "770-iRzNx", "771-bijsi", "772-zXnnw", "773-JxUka", "774-XGtzd", "775-ZyBLE", "776-EtAVt", "777-SEkwJ", "778-SqxBD", "779-Mdlrc", "780-USzRC", "781-acSyQ", "782-UrVpN", "783-IWKFY", "784-TTfDZ", "785-ABFzE", "786-qmRQi", "787-tPIdo", "788-KZwbf", "789-sfUgi", "790-zRSAB", "791-ObTdo", "792-vonRy", "793-pnrav", "794-VVELq", "795-denDC", "796-eaHgs", "797-CjfEN", "798-WpzAa", "799-SbhNc", "800-jONDo",
		"801-BlxHP", "802-lPJYe", "803-AGJOu", "804-gBSBw", "805-XXRkg", "806-JyriT", "807-GbwTp", "808-zdvQW", "809-bmxog", "810-OKdly", "811-rwmBb", "812-xArIE", "813-qdqyM", "814-UpfdQ", "815-XGHJm", "816-KGRMl", "817-EIHUS", "818-NASYs", "819-eEZUe", "820-vhgPG", "821-IGyWF", "822-dEuDB", "823-AUFJz", "824-eloCb", "825-tgzxj", "826-Rusid", "827-ecrZF", "828-Yuoac", "829-fnRfR", "830-MTbvU", "831-sHuyQ", "832-hwMmp", "833-abdIJ", "834-aUoTe", "835-XULZy", "836-fvNdg", "837-VRNXS", "838-oWGuY", "839-zrujF", "840-WHaSW", "841-rsORP", "842-lCtOj", "843-tFmCs", "844-oYmyx", "845-PtmYh", "846-OHmmU", "847-Aloff", "848-UAHjJ", "849-ITLQS", "850-FPgIp", "851-vfbao", "852-qYCYf", "853-LddGr", "854-xLIhQ", "855-CgeOC", "856-HywUC", "857-ItMPk", "858-NunDk", "859-gksTQ", "860-pEXaZ", "861-xVbkk", "862-ylwlv", "863-rpLqi", "864-pDdVi", "865-fgzzR", "866-YKckb", "867-AJxhe", "868-qLzuf", "869-Gzahl", "870-fWGYw", "871-IFvMi", "872-DvlWe", "873-SntcV", "874-yMgls", "875-meSxZ", "876-YvsBS", "877-gBcYb", "878-KeLEt", "879-NuXZV", "880-KkZZb", "881-zlMEA", "882-XkBWr", "883-tkpeh", "884-cPQWe", "885-YrGDa", "886-BaLGu", "887-wQLQP", "888-PqiwU", "889-WwVQx", "890-tJrcP", "891-OVmTH", "892-dOdSa", "893-gluLV", "894-KzjVN", "895-ezyqV", "896-QyyOW", "897-tHQJM", "898-bHiFd", "899-dcvQW", "900-NWGpm",
		"901-gZYYa", "902-gxeFk", "903-yWrgb", "904-wLFCc", "905-WphBg", "906-yEdTS", "907-DuWxF", "908-hinnh", "909-DxxHq", "910-gKREz", "911-zcSvw", "912-isGtQ", "913-UfTCL", "914-HbhWp", "915-kwctn", "916-mYxZe", "917-xcsUa", "918-lCSDl", "919-QVJEN", "920-dyjwR", "921-ijAaq", "922-HdKrz", "923-GtuEZ", "924-lRGJB", "925-cYJpb", "926-fBIJT", "927-XcydD", "928-FpZxd", "929-ifRXI", "930-NYBDm", "931-TeGKe", "932-VVmAk", "933-ekGzd", "934-wOVrS", "935-doeWH", "936-WtAvx", "937-qVdvG", "938-cdGqX", "939-HeTOm", "940-Rdref", "941-YQMIH", "942-lQHRP", "943-KYFqq", "944-UqMbu", "945-BDFUE", "946-NrSZx", "947-QKoCE", "948-PQQbQ", "949-dKqcf", "950-hNzAW", "951-OiiFT", "952-ndWyl", "953-ihUaZ", "954-GkRGh", "955-vCVEa", "956-PvtoQ", "957-oDiAH", "958-uFlXs", "959-irCFR", "960-wZumj", "961-mrcqT", "962-mtXUp", "963-zZYHf", "964-lfPGg", "965-ZXkKG", "966-CjXev", "967-EnFvu", "968-KZcLn", "969-dluBn", "970-qJweI", "971-azeyD", "972-eRefE", "973-Ywfiy", "974-ZMSij", "975-aAofs", "976-SpASD", "977-ThGoL", "978-GFmyh", "979-HMysM", "980-LHqTX", "981-UZdXt", "982-Zqnvz", "983-RQPnO", "984-eTCBC", "985-uXnAU", "986-SNtNM", "987-zIQge", "988-pneoe", "989-oMXQm", "990-aSIui", "991-LxFEn", "992-BYVXY", "993-yvbSV", "994-LwNHT", "995-bfhae", "996-yPnWk", "997-HoKGK", "998-yYEsU", "999-eAyJV", "1000-Lusnf",
	}

	return claims, scopes
}

func Test_isScopesValid(t *testing.T) {
	type args struct {
		claims claims
		scopes []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "scope valid",
			args: args{
				claims: claims{
					Authorization: authorization{
						Permissions: []permission{
							{Scopes: []string{"foo", "bar", "baz"}},
							{Scopes: []string{"qux", "fred"}},
						},
					},
				},
				scopes: []string{"bar", "fred"},
			},
			want: true,
		},
		{
			name: "scope valid 2",
			args: args{
				claims: claims{
					Authorization: authorization{
						Permissions: []permission{
							{Scopes: []string{"foo", "bar", "baz"}},
							{Scopes: []string{"qux", "fred"}},
						},
					},
				},
				scopes: []string{"qux", "thud"},
			},
			want: true,
		},
		{
			name: "scope valid 2",
			args: args{
				claims: claims{
					Authorization: authorization{
						Permissions: []permission{
							{Scopes: []string{"foo", "bar", "baz"}},
							{Scopes: []string{"qux", "fred"}},
						},
					},
				},
				scopes: []string{"thud", "chips"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScopesValid(tt.args.claims, tt.args.scopes); got != tt.want {
				t.Errorf("isScopesValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
