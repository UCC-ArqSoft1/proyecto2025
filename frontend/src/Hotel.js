import React from 'react';
import './Hotel.css';

function Hotel() {
  const hotel = {
    name: 'Hotel Cordoba Palace',
    location: 'Córdoba, Argentina',
    rating: 4.6,
    reviews: 324,
    pricePerNight: 105,
    imageUrl: 'https://images.unsplash.com/photo-1600585154340-be6161a56a0c?auto=format&fit=crop&w=1500&q=80',
    description: 'Con una ubicación privilegiada y servicios premium, el Hotel Cordoba Palace es ideal tanto para turistas como viajeros de negocios.',
    amenities: ['WiFi gratis', 'Pileta al aire libre', 'Estacionamiento', 'Desayuno incluido', 'Gimnasio', 'Pet friendly']
  };

  return (
    <div className="hotel-container">
      <img src={hotel.imageUrl} alt={hotel.name} className="hotel-image" />

      <div className="hotel-body">
        <div className="hotel-info">
          <h1>{hotel.name}</h1>
          <p className="location">{hotel.location}</p>
          <div className="rating">
            ⭐ {hotel.rating} · {hotel.reviews} opiniones
          </div>
          <p className="description">{hotel.description}</p>

          <h4>Servicios</h4>
          <ul className="amenities">
            {hotel.amenities.map((item, index) => (
              <li key={index}>✅ {item}</li>
            ))}
          </ul>
        </div>

        <div className="reservation-card">
          <div className="price">
            <span>${hotel.pricePerNight}</span> / noche
          </div>
          <button className="reserve-btn">Reservar</button>
        </div>
      </div>
    </div>
  );
}

export default Hotel;
