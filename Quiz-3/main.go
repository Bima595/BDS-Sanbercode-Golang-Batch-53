// main.go

package main

import (
	"Quiz-3/config"
	"Quiz-3/models"
	"Quiz-3/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Shape interface {
	Calculate() float64
}

type Datar interface {
	Shape
	Area() float64
	Perimeter() float64
}

type Ruang interface {
	Shape
	Volume() float64
	SurfaceArea() float64
}

type Persegi struct {
	Sisi float64
}

type PersegiPanjang struct {
	Panjang float64
	Lebar   float64
}

type Lingkaran struct {
	JariJari float64
}

type Kubus struct {
	Sisi float64
}

type Balok struct {
	Panjang float64
	Lebar   float64
	Tinggi  float64
}

type Tabung struct {
	JariJari float64
	Tinggi   float64
}

func (p Persegi) Area() float64 {
	return p.Sisi * p.Sisi
}

func (p Persegi) Perimeter() float64 {
	return 4 * p.Sisi
}

func (p PersegiPanjang) Area() float64 {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Perimeter() float64 {
	return 2*(p.Panjang+p.Lebar)
}

func (l Lingkaran) Area() float64 {
	return 3.14 * l.JariJari * l.JariJari
}

func (l Lingkaran) Perimeter() float64 {
	return 2 * 3.14 * l.JariJari
}

func (k Kubus) Volume() float64 {
	return k.Sisi * k.Sisi * k.Sisi
}

func (k Kubus) SurfaceArea() float64 {
	return 6 * k.Sisi * k.Sisi
}

func (b Balok) Volume() float64 {
	return b.Panjang * b.Lebar * b.Tinggi
}

func (b Balok) SurfaceArea() float64 {
	return 2 * (b.Panjang*b.Lebar + b.Panjang*b.Tinggi + b.Lebar*b.Tinggi)
}

func (t Tabung) Volume() float64 {
	return 3.14 * t.JariJari * t.JariJari * t.Tinggi
}

func (t Tabung) SurfaceArea() float64 {
	return 2 * 3.14 * t.JariJari * (t.JariJari + t.Tinggi)
}

func (p Persegi) Calculate() float64 {
	return 0
}

func (p PersegiPanjang) Calculate() float64 {
	return 0 
}

func (l Lingkaran) Calculate() float64 {
	return 0 
}

func (k Kubus) Calculate() float64 {
	return 0 
}

func (b Balok) Calculate() float64 {
	return 0 
}

func (t Tabung) Calculate() float64 {
	return 0 
}


func HitungHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	jenisBangun := params["jenisBangun"]
	jenisHitung := r.URL.Query().Get("hitung")

	var shape Shape
	var hasil float64

	switch jenisBangun {
	case "persegi":
		sisi, err := strconv.ParseFloat(r.URL.Query().Get("sisi"), 64)
		if err != nil {
			http.Error(w, "Invalid parameter: sisi", http.StatusBadRequest)
			return
		}
		shape = Persegi{Sisi: sisi}
	case "persegi-panjang":
		panjang, err := strconv.ParseFloat(r.URL.Query().Get("panjang"), 64)
		if err != nil {
			http.Error(w, "Invalid parameter: panjang", http.StatusBadRequest)
			return
		}
		lebar, err := strconv.ParseFloat(r.URL.Query().Get("lebar"), 64)
		if err != nil {
			http.Error(w, "Invalid parameter: lebar", http.StatusBadRequest)
			return
		}
		shape = PersegiPanjang{Panjang: panjang, Lebar: lebar}
	case "lingkaran":
		jariJari, err := strconv.ParseFloat(r.URL.Query().Get("jariJari"), 64)
		if err != nil {
			http.Error(w, "Invalid parameter: jariJari", http.StatusBadRequest)
			return
		}
		shape = Lingkaran{JariJari: jariJari}
	case "kubus":
		sisi, err := strconv.ParseFloat(r.URL.Query().Get("sisi"), 64)
		if err != nil {
			http.Error(w, "Invalid parameter: sisi", http.StatusBadRequest)
			return
		}
		shape = Kubus{Sisi: sisi}
	case "balok":
		panjang, err := strconv.ParseFloat(r.URL.Query().Get("panjang"), 64)
		if err != nil {
			http.Error(w, "Invalid parameter: panjang", http.StatusBadRequest)
			return
		}
		lebar, err := strconv.ParseFloat(r.URL.Query().Get("lebar"), 64)
		if err != nil {
			http.Error(w, "Invalid parameter: lebar", http.StatusBadRequest)
			return
		}
		tinggi, err := strconv.ParseFloat(r.URL.Query().Get("tinggi"), 64)
		if err != nil {
			http.Error(w, "Invalid parameter: tinggi", http.StatusBadRequest)
			return
		}
		shape = Balok{Panjang: panjang, Lebar: lebar, Tinggi: tinggi}
	case "tabung":
		jariJari, err := strconv.ParseFloat(r.URL.Query().Get("jariJari"), 64)
		if err != nil {
			http.Error(w, "Invalid parameter: jariJari", http.StatusBadRequest)
			return
		}
		tinggi, err := strconv.ParseFloat(r.URL.Query().Get("tinggi"), 64)
		if err != nil {
			http.Error(w, "Invalid parameter: tinggi", http.StatusBadRequest)
			return
		}
		shape = Tabung{JariJari: jariJari, Tinggi: tinggi}
	default:
		http.Error(w, "Invalid jenisBangun", http.StatusBadRequest)
		return
	}

	switch jenisHitung {
	case "luas":
		hasil = shape.(Datar).Area()
	case "keliling":
		hasil = shape.(Datar).Perimeter()
	case "volume":
		if ruang, ok := shape.(Ruang); ok {
			hasil = ruang.Volume()
		} else {
			http.Error(w, "Invalid jenisHitung for bangun datar", http.StatusBadRequest)
			return
		}
	case "luasPermukaan":
		if ruang, ok := shape.(Ruang); ok {
			hasil = ruang.SurfaceArea()
		} else {
			http.Error(w, "Invalid jenisHitung for bangun datar", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Invalid jenisHitung", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hasil perhitungan %s %s: %.2f", jenisBangun, jenisHitung, hasil)
}

var db *sql.DB
// Middleware
func basicAuthMiddleware(validCredentials map[string]string) mux.MiddlewareFunc {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if requiresAuth(r.Method) {
                user, pass, ok := r.BasicAuth()

                if !ok || !isValidCredentials(user, pass, validCredentials) {
                    w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
                    w.WriteHeader(http.StatusUnauthorized)
                    fmt.Fprint(w, "Unauthorized")
                    return
                }
            }

            next.ServeHTTP(w, r)
        })
    }
}

func requiresAuth(method string) bool {
    return method == "POST" || method == "PUT" || method == "DELETE"
}

func isValidCredentials(username, password string, validCredentials map[string]string) bool {
    storedPassword, ok := validCredentials[username]
    return ok && storedPassword == password
}


func main() {
	var err error
	db, err = config.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	
	basicAuthCredentials := map[string]string{
		"admin":   "password",
		"editor":  "secret",
		"trainer": "rahasia",
	}

	authMiddleware := basicAuthMiddleware(basicAuthCredentials)

	router := mux.NewRouter()

	router.HandleFunc("/bangun-datar/{jenisBangun}", HitungHandler).Methods("GET")
	router.HandleFunc("/bangun-ruang/{jenisBangun}", HitungHandler).Methods("GET")

	router.HandleFunc("/articles", GetArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", GetArticle).Methods("GET")
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.Use(authMiddleware)
	authRouter.HandleFunc("/articles", CreateArticle).Methods("POST")
	authRouter.HandleFunc("/articles/{id}", UpdateArticle).Methods("PUT")
	authRouter.HandleFunc("/articles/{id}", DeleteArticle).Methods("DELETE")

	router.HandleFunc("/categories", GetCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", GetCategory).Methods("GET")
	authRouterCategories := router.PathPrefix("/auth/categories").Subrouter()
	authRouterCategories.Use(authMiddleware)
	authRouterCategories.HandleFunc("/categories", CreateCategory).Methods("POST")
	authRouterCategories.HandleFunc("/categories/{id}", UpdateCategory).Methods("PUT")
	authRouterCategories.HandleFunc("/categories/{id}", DeleteCategory).Methods("DELETE")

	fmt.Println("Server Running at Port 3301")
	log.Fatal(http.ListenAndServe(":3301", router))
}

func getArticleLength(content string) string {
    contentLength := len(content)

    if contentLength <= 100 || utf8.RuneCountInString(content) <= 400 {
        return "pendek"
    } else if contentLength <= 200 || utf8.RuneCountInString(content) <= 800 {
        return "sedang"
    } else {
        return "panjang"
    }
}


func GetArticles(w http.ResponseWriter, r *http.Request) {
    articles := []models.Article{}

    query := "SELECT id, title, content, image_url, article_length, IFNULL(created_at, '0000-00-00 00:00:00') as created_at, IFNULL(update_at, '0000-00-00 00:00:00') as updated_at, category_id FROM article"
    
    queryParams := r.URL.Query()
    title := queryParams.Get("title")
    minYear := queryParams.Get("minYear")
    maxYear := queryParams.Get("maxYear")
    minWord := queryParams.Get("minWord")
    maxWord := queryParams.Get("maxWord")
    sortByTitle := queryParams.Get("sortByTitle")

    // Add filters
    if title != "" {
        query += fmt.Sprintf(" WHERE title LIKE '%%%s%%'", title)
    }
    if minYear != "" {
        if title != "" {
            query += " AND"
        } else {
            query += " WHERE"
        }
        query += fmt.Sprintf(" YEAR(created_at) >= %s", minYear)
    }
    if maxYear != "" {
        if title != "" || minYear != "" {
            query += " AND"
        } else {
            query += " WHERE"
        }
        query += fmt.Sprintf(" YEAR(created_at) <= %s", maxYear)
    }
    if minWord != "" {
        if title != "" || minYear != "" || maxYear != "" {
            query += " AND"
        } else {
            query += " WHERE"
        }
        query += fmt.Sprintf(" LENGTH(content) >= %s", minWord)
    }
    if maxWord != "" {
        if title != "" || minYear != "" || maxYear != "" || minWord != "" {
            query += " AND"
        } else {
            query += " WHERE"
        }
        query += fmt.Sprintf(" LENGTH(content) <= %s", maxWord)
    }

    if sortByTitle == "asc" {
        query += " ORDER BY title ASC"
    } else if sortByTitle == "desc" {
        query += " ORDER BY title DESC"
    }

    rows, err := db.Query(query)
    if err != nil {
        errorMessage := fmt.Sprintf("Error executing query: %v", err)
        utils.SendErrorResponse(w, http.StatusInternalServerError, errorMessage)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var article models.Article
        var createdAt, updatedAt string

        err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.ImageURL, &article.ArticleLength, &createdAt, &updatedAt, &article.CategoryID)
        if err != nil {
            errorMessage := fmt.Sprintf("Error scanning rows: %v", err)
            utils.SendErrorResponse(w, http.StatusInternalServerError, errorMessage)
            return
        }

        article.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
        article.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

        articles = append(articles, article)
    }

    utils.SendJSONResponse(w, http.StatusOK, articles)
}




func GetArticle(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    var article models.Article
    var createdAt, updatedAt string

    err := db.QueryRow("SELECT id, title, content, image_url, article_length, IFNULL(created_at, '0000-00-00 00:00:00') as created_at, IFNULL(update_at, '0000-00-00 00:00:00') as updated_at, category_id FROM article WHERE id = ?", id).Scan(
        &article.ID, &article.Title, &article.Content, &article.ImageURL, &article.ArticleLength, &createdAt, &updatedAt, &article.CategoryID,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            utils.SendErrorResponse(w, http.StatusNotFound, "Article not found")
        } else {
            errorMessage := fmt.Sprintf("Error getting article: %v", err)
            utils.SendErrorResponse(w, http.StatusInternalServerError, errorMessage)
        }
        return
    }

    article.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
    article.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

    utils.SendJSONResponse(w, http.StatusOK, article)
}


func CreateArticle(w http.ResponseWriter, r *http.Request) {
    var article models.Article

    err := json.NewDecoder(r.Body).Decode(&article)
    if err != nil {
        utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    currentTime := time.Now()

    articleLength := getArticleLength(article.Content)

    result, err := db.Exec("INSERT INTO article (title, content, image_url, article_length, created_at, update_at, category_id) VALUES (?, ?, ?, ?, ?, ?, ?)", article.Title, article.Content, article.ImageURL, articleLength, currentTime, currentTime, article.CategoryID)
    if err != nil {
        utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }

    id, _ := result.LastInsertId()
    article.ID = int(id)
    article.ArticleLength = articleLength

    utils.SendJSONResponse(w, http.StatusCreated, article)
}


func UpdateArticle(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    var article models.Article

    err := json.NewDecoder(r.Body).Decode(&article)
    if err != nil {
        utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    articleLength := getArticleLength(article.Content)

    result, err := db.Exec("UPDATE article SET title=?, content=?, image_url=?, article_length=?, updated_at=?, category_id=? WHERE id=?", article.Title, article.Content, article.ImageURL, articleLength, time.Now(), article.CategoryID, id)
    if err != nil {
        utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
        return
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        utils.SendErrorResponse(w, http.StatusNotFound, "Article not found")
        return
    }

    article.ID, _ = strconv.Atoi(id)
    article.ArticleLength = articleLength

    utils.SendJSONResponse(w, http.StatusOK, article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	result, err := db.Exec("DELETE FROM article WHERE id=?", id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		utils.SendErrorResponse(w, http.StatusNotFound, "Article not found")
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, map[string]string{"message": "Article deleted successfully"})
}


func GetCategories(w http.ResponseWriter, r *http.Request) {
    categories := []models.Category{}

    rows, err := db.Query("SELECT * FROM category")
    if err != nil {
        errorMessage := fmt.Sprintf("Error executing query: %v", err)
        utils.SendErrorResponse(w, http.StatusInternalServerError, errorMessage)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var category models.Category
        var createdAt, updatedAt mysql.NullTime

        err := rows.Scan(&category.ID, &category.Name, &createdAt, &updatedAt)
        if err != nil {
            errorMessage := fmt.Sprintf("Error scanning rows: %v", err)
            utils.SendErrorResponse(w, http.StatusInternalServerError, errorMessage)
            return
        }

        category.CreatedAt = createdAt.Time
        category.UpdatedAt = updatedAt.Time

        categories = append(categories, category)
    }

    utils.SendJSONResponse(w, http.StatusOK, categories)
}


func GetCategory(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    var category models.Category
    var createdAt, updatedAt mysql.NullTime

    err := db.QueryRow("SELECT * FROM category WHERE id = ?", id).Scan(&category.ID, &category.Name, &createdAt, &updatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            errorMessage := fmt.Sprintf("Category with ID %s not found", id)
            utils.SendErrorResponse(w, http.StatusNotFound, errorMessage)
            return
        }

        errorMessage := fmt.Sprintf("Error executing query: %v", err)
        utils.SendErrorResponse(w, http.StatusInternalServerError, errorMessage)
        return
    }

    category.CreatedAt = createdAt.Time
    category.UpdatedAt = updatedAt.Time

    utils.SendJSONResponse(w, http.StatusOK, category)
}


func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result, err := db.Exec("INSERT INTO category (name, created_at, updated_at) VALUES (?, ?, ?)", category.Name, time.Now(), time.Now())
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	id, _ := result.LastInsertId()
	category.ID = int(id)

	utils.SendJSONResponse(w, http.StatusCreated, category)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result, err := db.Exec("UPDATE category SET name=?, updated_at=? WHERE id=?", category.Name, time.Now(), id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		utils.SendErrorResponse(w, http.StatusNotFound, "Category not found")
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, category)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	result, err := db.Exec("DELETE FROM category WHERE id=?", id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		utils.SendErrorResponse(w, http.StatusNotFound, "Category not found")
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, map[string]string{"message": "Category deleted successfully"})
}

