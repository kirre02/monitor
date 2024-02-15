import { z } from "zod"

const URLSchema = z.string(). url()

function isValidURL(url: string): boolean {
    try {
        URLSchema.parse(url);
        return true; 
    } catch (error) {
        return false; 
    }
}

export default isValidURL