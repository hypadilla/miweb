package services_cache

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func init() {
	// inicializa la caché con un tiempo de vida de 5 minutos
	c = cache.New(5*time.Minute, 10*time.Minute)
}

func ObtenerDatosDeCache(clave string) ([]interface{}, error) {
    // intenta obtener los datos de la caché
    if datos, encontrado := c.Get(clave); encontrado {
        // los datos se encontraron en la caché, los devolvemos
        return datos.([]interface{}), nil
    }

    // los datos no se encontraron en la caché
    return nil, fmt.Errorf("los datos no se encontraron en la caché")
}

func GuardarDatosEnCache(clave string, datos []interface{}) {
    // almacena los datos en la caché
    c.Set(clave, datos, cache.DefaultExpiration)
}
