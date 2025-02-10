interface Address {
  id: number;
  street: string;
  city: string;
  state: string;
  zipcode: string;
  user_id: number;
  created_at: string;
  updated_at: string;
}

export interface User {
  id: number;
  email: string;
  first_name: string;
  last_name: string;
  username: string;
  phone: string;
  address: Address;
  created_at: string;
  updated_at: string;
}
