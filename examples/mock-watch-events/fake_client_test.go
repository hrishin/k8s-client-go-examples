package main

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func Test_get_pod_using_face_client(t *testing.T) {
	client := fake.NewSimpleClientset(&v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake",
			Namespace: "fake",
		},
	})

	t.Log("Fetch the pod by pod name using the client-go API ")
	{
		pod, err := client.CoreV1().Pods("fake").Get("fake", metav1.GetOptions{})
		t.Log("\tTest 0: checking the error code response")
		{
			if err != nil {
				t.Fatalf("\t%s\t Get pod returned an nexpected error : %v", failed, err)
			}
			t.Logf("\t%s\tClient go has return no error.", succeed)
		}

		t.Log("\tTest 1: verifying the retrived pod from client-go get pod API")
		{
			if pod.Name != "fake" {
				t.Fatalf("\t%s\t Got pod %s, expected %s", failed, pod.Name, "fake")
			}
			t.Logf("\t%s\tClient go has returned the expected pod", succeed)
		}
	}
}
