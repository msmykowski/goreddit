package reddit_test

import (
	. "github.com/msmykowski/goreddit/reddit"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("RedditPost", func() {
	var (
		redditPost RedditPost
		server     *ghttp.Server
	)

	Describe("GetImgurId", func() {
		Context("Reddit URL has no extension", func() {

			BeforeEach(func() {
				redditPost = RedditPost{
					Title: "Test",
					URL:   "http://imgur.com/txyQWYL",
					Score: 5,
				}
			})

			It("returns the Imgur ID", func() {
				imgurId := redditPost.GetImgurId()
				Expect(imgurId).To(Equal("txyQWYL"))
			})
		})

		Context("Reddit URL has an extension", func() {

			BeforeEach(func() {
				redditPost = RedditPost{
					Title: "Test",
					URL:   "http://imgur.com/txyQWYL.jpg",
					Score: 5,
				}
			})

			It("returns the Imgur ID", func() {
				imgurId := redditPost.GetImgurId()
				Expect(imgurId).To(Equal("txyQWYL"))
			})
		})

		Context("Reddit URL is an album indicated by /a/", func() {

			BeforeEach(func() {
				redditPost = RedditPost{
					Title: "Test",
					URL:   "http://imgur.com/a/txyQWYL",
					Score: 5,
				}
			})

			It("returns the Imgur ID", func() {
				imgurId := redditPost.GetImgurId()
				Expect(imgurId).To(Equal("txyQWYL"))
			})
		})

		Context("Reddit URL is a gallery indicated by /gallery/", func() {

			BeforeEach(func() {
				redditPost = RedditPost{
					Title: "Test",
					URL:   "http://imgur.com/gallery/txyQWYL",
					Score: 5,
				}
			})

			It("returns the Imgur ID", func() {
				imgurId := redditPost.GetImgurId()
				Expect(imgurId).To(Equal("txyQWYL"))
			})
		})
	})

	Describe("GetImgurUrl", func() {
		BeforeEach(func() {
			server = ghttp.NewServer()
		})

		AfterEach(func() {
			server.Close()
		})

		It("returns the Imgur URL", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "https://api.imgur.com/3/image/txyQWYL"),
					ghttp.RespondWith(200, "[]"),
				),
			)

			redditPost.GetImgurUrl("txyQWYL")
			Expect(server.ReceivedRequests()).Should(HaveLen(1))
		})
	})
})
