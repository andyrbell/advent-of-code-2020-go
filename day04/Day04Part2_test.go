package day04

import (
	"advent-of-code-2020-go/utils"
	"fmt"
	"testing"
)

func TestPart2Example(t *testing.T) {
	input := []string{
		"eyr:1972 cid:100",
		"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
		"",
		"iyr:2019",
		"hcl:#602927 eyr:1967 hgt:170cm",
		"ecl:grn pid:012533040 byr:1946",
		"",
		"hcl:dab227 iyr:2012",
		"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
		"",
		"hgt:59cm ecl:zzz",
		"eyr:2038 hcl:74454a iyr:2023",
		"pid:3556412378 byr:2007",
		"",
		"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
		"hcl:#623a2f",
		"",
		"eyr:2029 ecl:blu cid:129 byr:1989",
		"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
		"",
		"hcl:#888785",
		"hgt:164cm byr:2001 iyr:2015 cid:88",
		"pid:545766238 ecl:hzl",
		"eyr:2022",
		"",
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
		"",
	}
	result := solvePart2(input)
	if result != 4 {
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	input := utils.ReadStrings("day04_input.txt")
	result := solvePart2(input)
	fmt.Printf("Part 2: %v\n", result)
}

func TestByr(t *testing.T) {
	passport := utils.Passport{PassportData: map[string]string{"byr": "2002"}}
	if !passport.ValidBirthYear() {
		t.Errorf("byr 2002 should be valid")
	}
	passport = utils.Passport{PassportData: map[string]string{"byr": "2003"}}
	if passport.ValidBirthYear() {
		t.Errorf("byr 2003 should be invalid")
	}
}

func TestHgt(t *testing.T) {
	passport := utils.Passport{PassportData: map[string]string{"hgt": "60in"}}
	if !passport.ValidHeight() {
		t.Errorf("hgt 60in should be valid")
	}
	passport = utils.Passport{PassportData: map[string]string{"hgt": "190cm"}}
	if !passport.ValidHeight() {
		t.Errorf("hgt 190cm should be valid")
	}
	passport = utils.Passport{PassportData: map[string]string{"hgt": "190in"}}
	if passport.ValidHeight() {
		t.Errorf("hgt 190in should be invalid")
	}
	passport = utils.Passport{PassportData: map[string]string{"hgt": "190"}}
	if passport.ValidHeight() {
		t.Errorf("hgt 190 should be invalid")
	}
}

func TestHcl(t *testing.T) {
	passport := utils.Passport{PassportData: map[string]string{"hcl": "#123abc"}}
	if !passport.ValidHairColour() {
		t.Errorf("hcl #123abc should be valid")
	}
	passport = utils.Passport{PassportData: map[string]string{"hcl": "#123abz"}}
	if passport.ValidHairColour() {
		t.Errorf("hcl #123abz should be invalid")
	}
	passport = utils.Passport{PassportData: map[string]string{"hcl": "123abc"}}
	if passport.ValidHairColour() {
		t.Errorf("hcl 123abc should be invalid")
	}
}

func TestEcl(t *testing.T) {
	passport := utils.Passport{PassportData: map[string]string{"ecl": "brn"}}
	if !passport.ValidEyeColour() {
		t.Errorf("ecl brn should be valid")
	}
	passport = utils.Passport{PassportData: map[string]string{"ecl": "wat"}}
	if passport.ValidEyeColour() {
		t.Errorf("ecl wat should be invalid")
	}
}

func TestPid(t *testing.T) {
	passport := utils.Passport{PassportData: map[string]string{"pid": "000000001"}}
	if !passport.ValidPassportId() {
		t.Errorf("pid 000000001 should be valid")
	}
	passport = utils.Passport{PassportData: map[string]string{"pid": "0123456789"}}
	if passport.ValidPassportId() {
		t.Errorf("pid 0123456789 should be invalid")
	}
}
