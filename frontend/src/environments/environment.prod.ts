export const environment = {
	production: true,
	isMockEnabled: true, // You have to switch this, when your real back-end is done
	authTokenKey: 'authce9d77b308c149d5992a80073637e4d5',
	api_url: 'https://donatuplasma.org:8000/api/register/',
	api_url_simple: 'https://donatuplasma.org:8000',
	api_url_match: 'https://donatuplasma.org:8001/api/matcher/'
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
