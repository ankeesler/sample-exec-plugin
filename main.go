package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	clientauthenticationv1beta1 "k8s.io/client-go/pkg/apis/clientauthentication/v1beta1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/auth/exec"
)

func performBusinessLogicWithExecCredential(ec runtime.Object) {
	// Fill in any business logic here that requires the use of the provided ExecCredential.Spec.
	// E.g., make decision whether this plugin should require information from the user on stdin.

	exampleExecCredentialBusinessLogic(ec)
}

func performBusinessLogicWithRESTConfig(rc *rest.Config) {
	// Fill in any business logic here that requires the use of a REST config.
	// E.g., make anonymous requests to an on-cluster endpoint.

	exampleRESTConfigBusinessLogic(rc)
}

func main() {
	printfln("start")
	printfln("KUBERNETES_EXEC_INFO: %q", os.Getenv("KUBERNETES_EXEC_INFO"))

	ec, rc, err := exec.LoadExecCredentialFromEnv()
	if err != nil {
		dief("load: %q", err.Error())
	}

	performBusinessLogicWithExecCredential(ec)
	performBusinessLogicWithRESTConfig(rc)

	data, err := json.Marshal(ec)
	if err != nil {
		dief("marshal: %q", err.Error())
	}
	printfln("marshal: %q", string(data))

	fmt.Println(string(data))
}

func dief(format string, a ...interface{}) {
	reallyPrintf("error: "+format+"\n", a...)

	// Exit with a non-zero exit code to indicate to client-go that this exec plugin failed to obtain
	// a credential.
	os.Exit(1)
}

func printfln(format string, a ...interface{}) {
	if os.Getenv("QUIET") != "true" {
		reallyPrintf(format+"\n", a...)
	}
}

func reallyPrintf(format string, a ...interface{}) {
	// Always print to stderr since stdout is used to communicate credentials to client-go.
	fmt.Fprintf(os.Stderr, "sample-exec-plugin> "+format, a...)
}

func exampleExecCredentialBusinessLogic(ec runtime.Object) {
	ecBeta, ok := ec.(*clientauthenticationv1beta1.ExecCredential)
	if !ok {
		dief("cast failed: %#v\n", ec)
	}

	reallyPrintf("enter token: ")
	token, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		dief("cannot read stdin: %q", err.Error())
	}
	token = strings.TrimSpace(token)

	ecBeta.Status = &clientauthenticationv1beta1.ExecCredentialStatus{
		Token: token,
	}
}

func exampleRESTConfigBusinessLogic(rc *rest.Config) {
	c, err := kubernetes.NewForConfig(rc)
	if err != nil {
		dief("cannot create kubernetes client: %q", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if _, err := c.CoreV1().ConfigMaps("kube-public").Get(ctx, "cluster-info", metav1.GetOptions{}); err != nil {
		fmt.Fprintln(os.Stderr, "cannot find cluster-info: "+err.Error())
	}
}
