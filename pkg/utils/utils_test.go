package utils_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	utils "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/utils"
)

var _ = Describe("GUID", func() {
	It("should parse and process GUIDs correctly", func() {
		guidStr := "00:01:02:03:04:05:06:08"
		nextGuidStr := "00:01:02:03:04:05:06:09"

		guid, err := utils.ParseGUID(guidStr)
		Expect(err).NotTo(HaveOccurred())

		Expect(guid.String()).To(Equal(guidStr))
		Expect((guid + 1).String()).To(Equal(nextGuidStr))
	})
})

func TestUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Utils Suite")
}
