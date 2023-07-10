import axios from "axios";

export namespace api {
  type Param = number | string;
  type Note = {
    id: number;
    title: string;
    body: string;
  }

  type PaginationResponse = {
    status: boolean;
    total: number;
    page: number;
    prev_page: boolean;
    next_page: boolean;
    data: Note[];
  }

  export async function newNote(): Promise<number | undefined> {
    const data = await axios.post("/user/save", { title: "新便签", body: "写点什么" });
    if (data.data.status) return data.data.id;
  }

  export async function getNotes(page: Param): Promise<PaginationResponse> {
    const data = await axios.post(`/user/list?page=${page}`);
    return data.data as PaginationResponse;
  }

  export async function getNoteById(id: Param): Promise<Note | undefined> {
    const data = await axios.get(`/user/get?id=${id}`);
    if (data.data.status) return data.data as Note;
  }

  export async function saveNoteById(id: Param, title: string, body: string): Promise<boolean> {
    const data = await axios.post(`/user/update?id=${id}`, { id, title, body });
    return data.data.status;
  }

  export async function deleteNoteById(id: Param): Promise<boolean> {
    const data = await axios.post(`/user/delete?id=${id}`);
    return data.data.status;
  }
}
