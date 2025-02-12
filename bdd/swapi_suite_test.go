package main

import (
	"encoding/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tidwall/gjson"
	"log"
	"testing"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Starships Suite")
}

var _ = Describe("Verify Starships endpoint", Ordered, func() {
	var root *Root

	BeforeAll(func() {
		resp, err := newRestyClient().R().Get(BaseURL)
		errHandleFatalf(Format, err)
		if err := json.Unmarshal(resp.Body(), &root); err != nil {
			log.Fatalf("Error parsing JSON: %v", err)
		}
	})

	Describe("[/films] returns a list of films", func() {
		var respBody string

		BeforeAll(func() {
			resp, err := newRestyClient().R().Get(root.Films)
			errHandleFatalf(Format, err)
			respBody = resp.String()
		})

		Context("when extracting films titles", func() {
			It("should ContainElements of known titles", func() {
				expectedArr := []string{"A New Hope", "The Empire Strikes Back", "Return of the Jedi", "The Phantom Menace", "Attack of the Clones", "Revenge of the Sith"}
				actualArr := gjson.Get(respBody, "results.#.title").Value().([]interface{})
				Expect(actualArr).Should(ContainElements(expectedArr))
			})
		})

		Context("when extracting film where [episode_id:3]", func() {
			It("should contain 13 elements", func() {
				actualArr := gjson.Get(respBody, "results.#(episode_id==3).planets").Value().([]any)
				Expect(actualArr).Should(HaveLen(13))
			})
		})
	})

	Describe("[/people] returns a list of people", func() {
		var respBody string

		Context("when searching name [Chewbacca]", func() {

			BeforeAll(func() {
				resp, err := newRestyClient().R().Get(root.People + "?search=Chewbacca")
				errHandleFatalf(Format, err)
				respBody = resp.String()
			})

			It("should have [birth_year] Equal to 200BBY", func() {
				expectedArr := "200BBY"
				actualBY := gjson.Get(respBody, "results.0.birth_year").String()
				Expect(actualBY).Should(Equal(expectedArr))
			})

			It("should have [4] films", func() {
				actualArr := gjson.Get(respBody, "results.0.films").Value().([]any)
				Expect(actualArr).Should(HaveLen(4))
			})
		})
	})
})
