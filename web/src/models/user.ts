interface Address {
  id: string;
  street: string;
  city: string;
  state: string;
  zipcode: string;
  user_id: string;
  created_at: string;
  updated_at: string;
}

export interface User {
  id: string;
  email: string;
  name: string;
  username: string;
  phone: string;
  address: Address;
  created_at: string;
  updated_at: string;
}
