import { Effect } from "./effect.model";

export class Config {
    public plugins: {
        id: string,
        name: string,
        filename: string
    }[];
    public effects: Effect[];
}
