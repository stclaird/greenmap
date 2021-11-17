{
    LocationPoint:
      { $near:
         {
           $geometry: { type: "Point",  coordinates: [ 50.8539678, -0.969052 ] },
           $minDistance: 1000,
           $maxDistance: 5000
         }
      }
}