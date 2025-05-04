package data

import (
	"errors"

	"fixable.com/fixable/internal/models"
	"gorm.io/gorm"
)

var categorias = []models.Categoria{
	{Nombre: "Ferreteria"},
	{Nombre: "Carpinteria"},
	{Nombre: "Electrica"},
	{Nombre: "Plomeria"},
}

var servicios = []models.Servicio{
	{Nombre: "Venta de herramientas", Imagen: "https://th.bing.com/th/id/R.455abe422abd1a99994bf677ad046425?rik=2EdSruzZVlAxLQ&pid=ImgRaw&r=0", CategoriaId: 1, Direccion: "Calle 123, Ciudad", Email: "ferreteria@gmail.com", Telefono: "825-1234", Descripcion: "Ofrecemos herramientas de calidad para todo tipo de trabajos de construcción y mantenimiento."},
	{Nombre: "Instalaciones eléctricas", Imagen: "https://th.bing.com/th/id/R.42c8a2387238c8de3d8adc1e32a41d5f?rik=8q0uk%2b9wGX9dNA&riu=http%3a%2f%2fwww.electricistasadomicilio.mx%2fwp-content%2fuploads%2f2020%2f07%2f%c2%a1Realiza-tu-pedido-Tel_-55-8887-2966-11.jpg&ehk=vJv3LJSXDZk6spdV%2fYLpftP9QOYyS%2bkNCfaWfZFLXpk%3d&risl=&pid=ImgRaw&r=0", CategoriaId: 3, Direccion: "Calle 456, Ciudad", Email: "electrica@gmail.com", Telefono: "550-5609", Descripcion: "Especialistas en instalaciones eléctricas residenciales, comerciales e industriales."},
	{Nombre: "Fabricación de muebles", Imagen: "https://cdn.5dias.com.py/uploads/Dise%C3%B1o_sin_t%C3%ADtulo_-_2022-09-12T084448.833.jpg", CategoriaId: 2, Direccion: "Calle 789, Ciudad", Email: "carpinteria@gmail.com", Telefono: "812-98076", Descripcion: "Diseñamos y fabricamos muebles personalizados para tu hogar o negocio."},
	{Nombre: "Reparación de tuberías", Imagen: "https://www.arqhys.com/wp-content/uploads/2012/12/Plomeria.jpg", CategoriaId: 4, Direccion: "Calle 321, Ciudad", Email: "plomeria@gmail.com", Telefono: "852-4321", Descripcion: "Reparamos tuberías, grifos y sistemas de agua con rapidez y eficiencia."},
	{Nombre: "Materiales de construcción", Imagen: "https://lirp.cdn-website.com/7765aaa7/dms3rep/multi/opt/b1-1920w.jpg", CategoriaId: 1, Direccion: "Calle 654, Ciudad", Email: "ferreteria2@gmail.com", Telefono: "651-6789", Descripcion: "Venta de materiales de construcción para proyectos grandes y pequeños."},
}

func SeedData(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Categoria{}); err == nil && db.Migrator().HasTable(&models.Categoria{}) {
		if err := db.First(&models.Categoria{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.Create(&categorias)
			if result.Error != nil {
				return result.Error
			}
		}
	}

	if err := db.AutoMigrate(&models.Servicio{}); err == nil && db.Migrator().HasTable(&models.Servicio{}) {
		if err := db.First(&models.Servicio{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.Create(&servicios)
			if result.Error != nil {
				return result.Error
			}
		}
	}
  db.AutoMigrate(&models.Comentario{})

	return nil

}
