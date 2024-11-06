import { User } from './userTypes'

export interface Group {
    id: number;
    name: string;
    members: User[];
    isOwner: boolean;
}
