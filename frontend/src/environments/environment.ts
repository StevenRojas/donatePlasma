// This file can be replaced during build by using the `fileReplacements` array.
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

export const environment = {
  production: false,
  isMockEnabled: true, // You have to switch this, when your real back-end is done
  authTokenKey: 'authce9d77b308c149d5992a80073637e4d5',
  api_url: 'http://localhost:8000/api/register/',
  api_url_simple: 'http://localhost:8000',
  api_url_match: 'http://localhost:8001/api/matcher/'
};


export const bloodTypes = [
  { value: 1, display: 'O-' },
  { value: 2, display: 'O+' },
  { value: 3, display: 'A-' },
  { value: 4, display: 'A+' },
  { value: 5, display: 'B-' },
  { value: 6, display: 'B+' },
  { value: 7, display: 'AB-' },
  { value: 8, display: 'AB+' }
];

export const cities =[
  { value: 1, display: 'Cochabamba' },
  { value: 2, display: 'La Paz' },
  { value: 3, display: 'Santa Cruz' },
  { value: 4, display: 'Oruro' },
  { value: 5, display: 'Potosi' },
  { value: 6, display: 'Tarija' },
  { value: 7, display: 'Chuquisaca' },
  { value: 8, display: 'Beni' },
  { value: 9, display: 'Pando' }
];
/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/dist/zone-error';  // Included with Angular CLI.
