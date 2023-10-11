import { Component, Input, OnInit } from '@angular/core';
import { Todo } from './_common/model/todo.model';
import { TodoService } from './_common/service/todo.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  todoList: Todo[] = [];
  @Input() todoTitle = '';
  isModalOpen = false;
  todoEditId = 0;

  constructor(private todoService: TodoService) {}

  ngOnInit(): void {
    this.getList();
  }

  openModal() {
    this.isModalOpen = true;
  }

  closeModal() {
    this.isModalOpen = false;
  }

  getList() {
    this.todoService.getTodo().subscribe((data) => {
      this.todoList = data;
    });
  }

  toggleCompleted(todo: Todo) {
    this.todoService.toggleCompleted(todo.id).subscribe((todo) => {
      this.getList();
    });
  }

  onSubmit() {
    if (this.todoEditId === 0) {
      this.todoService.createTodo(this.todoTitle).subscribe((todo) => {
        this.getList();
        this.todoTitle = '';
        this.closeModal();
      });
    } else {
      this.todoService
        .updateTodoTitle(this.todoEditId, this.todoTitle)
        .subscribe((todo) => {
          this.getList();
          this.todoTitle = '';
          this.todoEditId = 0;
          this.closeModal();
        });
    }
  }

  onEditClick(todo: Todo) {
    this.todoEditId = todo.id;
    this.todoTitle = todo.title;
    this.openModal();
  }

  onDeleteClick(id: number) {
    this.todoService.deleteTodoById(id).subscribe(() => {
      this.getList();
    });
  }
}
