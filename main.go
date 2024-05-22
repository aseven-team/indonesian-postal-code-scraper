package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Province struct {
	Code string
	Name string
}

type City struct {
	Code     string
	Name     string
	Province string
}

type District struct {
	Code string
	Name string
	City string
}

type Village struct {
	Code     string
	Name     string
	District string
}

func main() {
	getProvinces()
	getCities()
	getDistricts()
	getVillages()
}

func getVillages() {
	fmt.Println("Scraping villages data...")

	os.Remove("output/villages.csv")

	totalPage := getTotalPage("https://kodepos.posindonesia.co.id/kelurahandesalist?recperpage=100")

	for i := 1; i <= totalPage; i++ {
		getVillagesByPage(i)
	}
}

func getVillagesByPage(page int) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Page", r.URL.String())
	})

	var villages []Village

	c.OnHTML("table#tbl_kelurahandesalist > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			village := Village{
				Code:     el.ChildText("td:nth-of-type(7)"),
				Name:     el.ChildText("td:nth-of-type(8)"),
				District: el.ChildText("td:nth-of-type(5)"),
			}

			villages = append(villages, village)
		})
	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.OpenFile("output/villages.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		if page == 1 {
			writer.Write([]string{"Kode", "Nama", "Kecamatan"})
		}

		for _, village := range villages {
			writer.Write([]string{village.Code, village.Name, village.District})
		}

		defer writer.Flush()
	})

	c.Visit(fmt.Sprintf("https://kodepos.posindonesia.co.id/kelurahandesalist?recperpage=100&page=%d", page))
}

func getDistricts() {
	fmt.Println("Scraping districts data...")

	os.Remove("output/districts.csv")

	totalPage := getTotalPage("https://kodepos.posindonesia.co.id/kecamatanlist?recperpage=100")

	for i := 1; i <= totalPage; i++ {
		getDistrictsByPage(i)
	}
}

func getDistrictsByPage(page int) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Page", r.URL.String())
	})

	var districts []District

	c.OnHTML("table#tbl_kecamatanlist > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			district := District{
				Code: el.ChildText("td:nth-of-type(5)"),
				Name: el.ChildText("td:nth-of-type(6)"),
				City: el.ChildText("td:nth-of-type(4)"),
			}

			districts = append(districts, district)
		})
	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.OpenFile("output/districts.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		if page == 1 {
			writer.Write([]string{"Kode", "Nama", "Kota"})
		}

		for _, district := range districts {
			writer.Write([]string{district.Code, district.Name, district.City})
		}

		defer writer.Flush()
	})

	c.Visit(fmt.Sprintf("https://kodepos.posindonesia.co.id/kecamatanlist?recperpage=100&page=%d", page))
}

func getCities() {
	fmt.Println("Scraping cities data...")

	os.Remove("output/cities.csv")

	totalPage := getTotalPage("https://kodepos.posindonesia.co.id/kabupatenkotalist?recperpage=100")

	for i := 1; i <= totalPage; i++ {
		getCitiesByPage(i)
	}
}

func getCitiesByPage(page int) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Page", r.URL.String())
	})

	var cities []City

	c.OnHTML("table#tbl_kabupatenkotalist > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			city := City{
				Code:     el.ChildText("td:nth-of-type(5)"),
				Name:     el.ChildText("td:nth-of-type(6)"),
				Province: el.ChildText("td:nth-of-type(3)"),
			}

			cities = append(cities, city)
		})
	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.OpenFile("output/cities.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		if page == 1 {
			writer.Write([]string{"Kode", "Nama", "Provinsi"})
		}

		for _, city := range cities {
			writer.Write([]string{city.Code, city.Name, city.Province})
		}

		defer writer.Flush()
	})

	c.Visit(fmt.Sprintf("https://kodepos.posindonesia.co.id/kabupatenkotalist?recperpage=100&page=%d", page))
}

func getProvinces() {
	fmt.Println("Scraping provinces data...")

	os.Remove("output/provinces.csv")

	totalPage := getTotalPage("https://kodepos.posindonesia.co.id/propinsilist?recperpage=100")

	for i := 1; i <= totalPage; i++ {
		getProvincesByPage(i)
	}
}

func getProvincesByPage(page int) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Page", r.URL.String())
	})

	var provinces []Province

	c.OnHTML("table#tbl_propinsilist > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			province := Province{
				Code: el.ChildText("td:nth-of-type(2)"),
				Name: el.ChildText("td:nth-of-type(3)"),
			}

			provinces = append(provinces, province)
		})
	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.OpenFile("output/provinces.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		if page == 1 {
			writer.Write([]string{"Kode", "Nama"})
		}

		for _, province := range provinces {
			writer.Write([]string{province.Code, province.Name})
		}

		defer writer.Flush()
	})

	c.Visit(fmt.Sprintf("https://kodepos.posindonesia.co.id/propinsilist?recperpage=100&page=%d", page))
}

func getTotalPage(url string) int {
	c := colly.NewCollector()

	var totalPage int

	c.OnHTML("div.ew-pager > span", func(e *colly.HTMLElement) {
		if totalPage != 0 {
			return
		}

		if strings.HasPrefix(strings.TrimSpace(e.Text), "dari") {
			pageText := strings.TrimSpace(strings.Replace(e.Text, "dari", "", 1))
			page, err := strconv.Atoi(pageText)

			if err != nil {
				fmt.Println("Failed to convert page number to integer", err)
			}

			totalPage = page
		}
	})

	c.Visit(url)

	return totalPage
}
