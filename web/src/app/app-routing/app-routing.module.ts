import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from '../login/login.component';
import { DashboardComponent } from '../dashboard/dashboard.component';
import { UserComponent } from '../user/user.component';
import { UsersComponent } from '../users/users.component';
import { SiteComponent } from '../site/site.component';
import { SitesComponent } from '../sites/sites.component';
import { RoleComponent } from '../role/role.component';
import { UserCreateComponent } from '../user-create/user-create.component';
import { SiteCreateComponent } from '../site/site-create/site-create.component';

/*
Routing list, can't use /api route for anything since the go server will catch it an use for itself.
 */
const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  { path: 'login', component: LoginComponent },
  { path: 'dashboard', component: DashboardComponent },
  { path: 'user/create', component: UserCreateComponent },
  { path: 'user/:id', component: UserComponent },
  { path: 'users', component: UsersComponent },
  { path: 'site/create', component: SiteCreateComponent },
  { path: 'site/:id', component: SiteComponent },
  { path: 'sites', component: SitesComponent },
  { path: 'role/:id', component: RoleComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
