import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MomentModule } from 'ngx-moment';
import { SharedModule } from '../shared.module';
import { SystemRoutingModule } from './system-routing.module';
import { SystemStatusComponent } from './status/systemstatus.component';
import { NbCardModule, NbListModule } from '@nebular/theme';
@NgModule({
  declarations: [SystemStatusComponent ],
  imports: [
    CommonModule,
    SystemRoutingModule,
    MomentModule,
    SharedModule,
    NbCardModule,
    NbListModule
  ]
})
export class SystemModule { }