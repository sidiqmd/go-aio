import { Component, OnInit } from '@angular/core';
import { Todo } from '../_common/model/todo.model';
import { TodoService } from '../_common/service/todo.service';

@Component({
  selector: 'app-todo-list',
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.css'],
})
export class TodoListComponent implements OnInit {
  todos: Todo[] = [];

  constructor(private todoService: TodoService) {}

  ngOnInit(): void {
    console.log('TodoListComponent.ngOnInit()');
    // this.todoService.getTodo().subscribe((todos) => {
    //   this.todos = todos;
    //   console.log(this.todos);
    // });
  }
}
