// requires script for bootstrapping google maps
let geocoder;
async function initGeocoder() {
  const { Geocoder } = await google.maps.importLibrary("geocoding");
  geocoder = new Geocoder();
}
initGeocoder();

async function isUK(lat, lng) {
  const g = await geocoder.geocode({ location: { lat, lng } });

  return g.results.some((r) => r.formatted_address === "United Kingdom");
}

async function addUKHexagons(hex, hexagons) {
  if (hexagons.includes(hex)) {
    return hexagons;
  }
  const [lat, lng] = h3.cellToLatLng(hex);
  const uk = await isUK(lat, lng);
  if (!uk) {
    return hexagons;
  }
  hexagons.push(hex);
  for (const neighbour of h3.gridDisk(hex, 1)) {
    hexagons = await addUKHexagons(neighbour, hexagons);
  }
  return hexagons;
}

export async function addHexagonsFromStartingPosition(lat, lng) {
  const startingHex = h3.latLngToCell(lat, lng, 5);
  return addUKHexagons(startingHex, []);
}
