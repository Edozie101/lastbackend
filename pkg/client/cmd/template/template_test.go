package template_test

import (
	"encoding/json"
	h "github.com/lastbackend/lastbackend/libs/http"
	"github.com/lastbackend/lastbackend/pkg/client/cmd/template"
	"github.com/lastbackend/lastbackend/pkg/client/context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {

	const token = "mocktoken"

	var (
		err error
		ctx = context.Mock()
	)

	ctx.Token = token

	//------------------------------------------------------------------------------------------
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tk := r.Header.Get("Authorization")
		assert.NotEmpty(t, tk, "token should be not empty")
		assert.Equal(t, tk, "Bearer "+token, "they should be equal")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}

		assert.Empty(t, body, "body should be empty")

		var temp = make(map[string][]string)

		temp["test_temp_1"] = []string{"ver. 1.1", "ver 2.2"}
		temp["test_temp_2"] = []string{"ver. 1.1", "ver 2.2", "ver. 3.3"}

		buff, err := json.Marshal(temp)

		if err != nil {
			t.Error(err)
			return
		}

		w.WriteHeader(200)
		_, err = w.Write(buff)
		if err != nil {
			t.Error(err)
			return
		}
	}))
	defer server.Close()
	//------------------------------------------------------------------------------------------

	ctx.HTTP = h.New(server.URL)

	_, err = template.List()

	if err != nil {
		t.Error(err)
		return
	}
}
