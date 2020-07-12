export class DonorModel {
    id: number;
    blood_type_id: number;
    name: string;
    cell: string;
    email: string;
    city_id: number;
    public: boolean;

    public getPrefix() {
        return 'donor';
    }
}
