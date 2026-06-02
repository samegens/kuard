package e2e_test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func imagePrefix() string {
	if p := os.Getenv("KUARD_IMAGE_PREFIX"); p != "" {
		return p
	}
	return "blauwelucht/kuard-amd64"
}

func startContainer(image string) (cid, baseURL string) {
	out, err := exec.Command("docker", "run", "-d", "-p", "0:8080", image).Output()
	Expect(err).NotTo(HaveOccurred(), "failed to start container for image %s", image)
	cid = strings.TrimSpace(string(out))

	portOut, err := exec.Command("docker", "port", cid, "8080").Output()
	Expect(err).NotTo(HaveOccurred())
	firstLine := strings.SplitN(strings.TrimSpace(string(portOut)), "\n", 2)[0]
	parts := strings.Split(firstLine, ":")
	port := parts[len(parts)-1]
	baseURL = fmt.Sprintf("http://localhost:%s", port)
	return
}

var _ = DescribeTable("kuard",
	func(fakever string) {
		image := fmt.Sprintf("%s:%s", imagePrefix(), fakever)

		cid, baseURL := startContainer(image)
		DeferCleanup(exec.Command("docker", "stop", cid).Run)

		By("waiting for the app to be ready")
		Eventually(func() error {
			resp, err := http.Get(baseURL + "/")
			if err != nil {
				return err
			}
			resp.Body.Close()
			return nil
		}, 10*time.Second, time.Second).Should(Succeed())

		By("serving index page with KUAR Demo title")
		resp, err := http.Get(baseURL + "/")
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		body, err := io.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(body)).To(ContainSubstring("KUAR Demo"))

		By("returning 200 on /healthy")
		resp, err = http.Get(baseURL + "/healthy")
		Expect(err).NotTo(HaveOccurred())
		resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))

		By("returning 200 on /ready")
		resp, err = http.Get(baseURL + "/ready")
		Expect(err).NotTo(HaveOccurred())
		resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
	},
	Entry("blue", "blue"),
	Entry("green", "green"),
	Entry("purple", "purple"),
)
