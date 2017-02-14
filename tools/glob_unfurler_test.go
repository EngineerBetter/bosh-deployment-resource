package tools_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"io/ioutil"
	"github.com/cloudfoundry/bosh-deployment-resource/tools"
)

var _ = Describe("GlobUnfurler", func() {
	var (
		releaseOne, releaseTwo, releaseThree *os.File
		releaseDir string
	)

	BeforeEach(func() {
		releaseDir, _ = ioutil.TempDir("", "primary-releases")

		releaseOne, _ = ioutil.TempFile(releaseDir, "release-one")
		releaseOne.Close()

		releaseTwo, _ = ioutil.TempFile(releaseDir, "release-two")
		releaseTwo.Close()

		releaseThree, _ = ioutil.TempFile(releaseDir, "coolio-three")
		releaseThree.Close()
	})

	It("returns all filepaths matching the globs", func() {
		filepaths, err := tools.UnfurlGlobs(
			releaseDir, []string{
				"release-*",
				"coolio-three*",
			},
		)

		Expect(err).ToNot(HaveOccurred())
		Expect(filepaths).To(ConsistOf(
			releaseOne.Name(),
			releaseTwo.Name(),
			releaseThree.Name(),
		))
	})

	Context("when some globs unfurl to the same file", func() {
		It("removes duplicate filepaths", func() {
			filepaths, err := tools.UnfurlGlobs(
				releaseDir, []string{
					"release-*",
					"rel*se-*",
				},
			)

			Expect(err).ToNot(HaveOccurred())
			Expect(filepaths).To(ConsistOf(
				releaseOne.Name(),
				releaseTwo.Name(),
			))
		})
	})

	Context("when a bad glob is passed", func() {
		It("returns an error", func() {
			_, err := tools.UnfurlGlobs(releaseDir, []string{"/["})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("/[ is not a valid file glob"))
		})
	})
})