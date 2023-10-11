import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Todo } from '../model/todo.model';

@Injectable({
  providedIn: 'root',
})
export class TodoService {
  private baseUrl = '/api/todo';

  constructor(private http: HttpClient) {}

  getTodo(): Observable<Todo[]> {
    return this.http.get<Todo[]>(this.baseUrl);
  }

  toggleCompleted(id: number): Observable<Todo> {
    return this.http.patch<Todo>(`${this.baseUrl}/${id}`, null);
  }

  createTodo(title: string): Observable<Todo> {
    return this.http.post<Todo>(this.baseUrl, { title });
  }

  updateTodoTitle(id: number, title: string): Observable<Todo> {
    return this.http.put<Todo>(`${this.baseUrl}/${id}`, { title });
  }

  deleteTodoById(id: number): Observable<void> {
    return this.http.delete<void>(`${this.baseUrl}/${id}`);
  }
}
