package api

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/google/uuid"
)

type response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

// errorInfo ...
type errorInfo struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func HandleInternalWithMessage(w http.ResponseWriter, err error, message string) error {
	if err == nil {
		return nil
	}

	log.Println(message+" ", err)
	w.WriteHeader(http.StatusInternalServerError)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusInternalServerError,
			Message: message,
		}})
	return err
}

func HandleBadRequestErrWithMessage(w http.ResponseWriter, err error, message string) error {
	if err == nil {
		return nil
	}

	log.Println(message+" ", err)
	w.WriteHeader(http.StatusBadRequest)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusBadRequest,
			Message: message + ": " + err.Error(),
		}})
	return err
}

func HandleUnauthorizedWithMessage(w http.ResponseWriter, message string) {
	w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
	w.WriteHeader(http.StatusUnauthorized)

	log.Println(message)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusUnauthorized,
			Message: message,
		}})
}

func HandleBadRequestResponse(w http.ResponseWriter, message string) {
	log.Println(message)
	w.WriteHeader(http.StatusBadRequest)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  http.StatusBadRequest,
			Message: message,
		}})
}

func HandleErrorResponse(w http.ResponseWriter, errCode int, errMessage string) {
	log.Println(errMessage+" code:", errCode)

	w.WriteHeader(errCode)
	writeJSON(w, response{Error: true,
		Data: errorInfo{
			Status:  errCode,
			Message: errMessage,
		}})
}

func parsePageQueryParam(r *http.Request) (int, error) {
	pageparam := r.URL.Query().Get("page")
	if pageparam == "" {
		return 1, nil
	}

	page, err := strconv.Atoi(pageparam)
	if err != nil {
		return 0, err
	}
	if page < 0 {
		return 0, errors.New("page must be an positive integer")
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func parseLimitQueryParam(r *http.Request) (int, error) {

	limitparam := r.URL.Query().Get("limit")
	if limitparam == "" {
		return 10, nil
	}

	page, err := strconv.Atoi(limitparam)
	if err != nil {
		return 0, err
	}
	if page < 0 {
		return 0, errors.New("limit must be an positive integer")
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func validateUUID(ID string) bool {
	_, err := uuid.Parse(ID)
	return err == nil
}

func BodyParser(r *http.Request, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(&body)
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	bytes, _ := json.MarshalIndent(data, "", "  ")

	w.Header().Set("Content-Type", "Application/json")
	w.Write(bytes)
}

func WriteJSONWithSuccess(w http.ResponseWriter, data interface{}) {
	data = response{
		Error: false,
		Data:  data,
	}
	bytes, _ := json.MarshalIndent(data, "", "  ")
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

// ParsePagination parses page and limit value from request query
func ParsePagination(r *http.Request) (page int, limit int, err error) {
	q := r.URL.Query()

	pageStr := q.Get("page")

	if pageStr == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return 0, 0, err
		}
		if page == 0 {
			page = 1
		}
	}

	limitStr := q.Get("limit")

	if limitStr == "" {
		limit = 0
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return 0, 0, err
		}
	}

	return page, limit, nil
}

// MarshalToJSONString ...
func MarshalToJSONString(data interface{}) string {

	bytes, _ := json.MarshalIndent(data, "", "  ")
	return string(bytes)
}

// ReadJSON reads JSON from http.Response and parses it into `out`
func ReadJSON(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()

	// if resp.StatusCode >= 400 {
	// 	body, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	// }

	if out == nil {
		io.Copy(ioutil.Discard, resp.Body)
		return nil
	}

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
