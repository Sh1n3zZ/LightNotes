import axios from "axios";
import { ref } from "vue";
import type { Ref } from "vue";

export namespace api {
  type Param = number | string;
  export type Note = {
    id: number;
    title: string;
    body: string;
    updated_at: string;
  }

  export type PaginationResponse = {
    status: boolean;
    total: number;
    page: number;
    prev_page: boolean;
    next_page: boolean;
    notes: Note[];
  }

  export async function newNote(): Promise<number | undefined> {
    const data = await axios.post("/user/save", { title: "新便签", body: "写点什么" });
    if (data.data.status) return data.data.id;
  }

  export async function getNotes(page: Param): Promise<PaginationResponse> {
    const data = await axios.get(`/user/list?page=${page}`);
    return data.data as PaginationResponse;
  }

  export async function getNoteById(id: Param): Promise<Note | undefined> {
    const data = await axios.get(`/user/get?id=${id}`);
    if (data.data.status) return data.data.note as Note;
  }

  export async function saveNoteById(id: Param, title: string, body: string): Promise<boolean> {
    const data = await axios.post(`/user/update?id=${id}`, { id, title, body });
    return data.data.status;
  }

  export async function deleteNoteById(id: Param): Promise<boolean> {
    const data = await axios.post(`/user/delete?id=${id}`);
    return data.data.status;
  }

  export function searchNotes(id: Param, data: Note[]): number {
    for (let i = 0; i < data.length; i++) {
      if (data[i].id == id) return i;
    }
    return -1;
  }

  export class NotePagination {
    public page: Ref<number>;
    public total: Ref<number>;
    protected data: Ref<Note[]>;
    protected end: boolean;

    public constructor() {
      this.page = ref(1);
      this.total = ref(0);
      this.data = ref([]);
      this.end = false;
    }

    public getRef(): Ref<Note[]> {
      return this.data;
    }

    public async update(): Promise<void> {
      if (this.end) return;
      const data = await getNotes(this.page.value);
      this.total.value = data.total;
      this.data.value = this.data.value.concat(data.notes);
      if (data.next_page) {
        this.page.value++;
      } else {
        this.end = true;
      }
    }

    public new(id: number): void {
      this.data.value.unshift({
        id,
        title: "新便签",
        body: "写点什么",
        updated_at: new Date(Date.now() + 8 * 3600000).toString(),
      });
    }

    public save(id: number, title: string, body: string): void {
      const index = searchNotes(id, this.data.value);
      if (index != -1) {
        this.data.value[index].title = title;
        this.data.value[index].body = body;
        this.data.value[index].updated_at = new Date(Date.now() + 8 * 3600000).toString();
      }
    }
  }
}
