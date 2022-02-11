import { MyNodeService } from './../mynodeservice';
import { Component, OnInit, ChangeDetectorRef, Input, OnDestroy, EventEmitter, Output } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { OTNodeSummaryModel } from './dataholders-models';
import { MomentModule } from 'ngx-moment';
import { HubHttpService } from '../../hub-http-service';
declare const $: any;
import Web3 from 'web3';
import { ServerDataSource } from 'ng2-smart-table';
import { DecimalPipe } from '@angular/common';
import { DataHolderDetailedModel } from '../dataholder/dataholder-models';
import { NbToastrService, NbGlobalLogicalPosition, NbToastrConfig, NbComponentStatus  } from '@nebular/theme';
import { MyNodeModel } from '../mynodemodel';
import { ServerSourceConf } from 'ng2-smart-table/lib/lib/data-source/server/server-source.conf';
import { OnlineIndicatorRenderComponent } from './onlineindicator.component';
@Component({
  selector: 'app-dataholders',
  templateUrl: './dataholders.component.html',
  styleUrls: ['./dataholders.component.scss']
})
export class DataHoldersComponent implements OnInit, OnDestroy {
  getNodesObserver: any;
  web3: any;
  settings: any;

  constructor(private http: HttpClient, private chRef: ChangeDetectorRef, private myNodeService: MyNodeService,
    private httpService: HubHttpService, private toastrService: NbToastrService) {
    this.isTableInit = false;
    this.isLoading = true;
    this.failedLoading = false;
    this.web3 = new Web3();
  }

  NodeModel: OTNodeSummaryModel[];
  dataTable: any;
  exportOptionsObj: any;
  isTableInit: boolean;
  failedLoading: boolean;
  isLoading: boolean;
  isDarkTheme: boolean;
  @Input() hideBreadcrumb: boolean;
  @Input() showOnlyMyNodes: string;
  @Input() managementWallet: string;
  @Output() afterLoadWithCount = new EventEmitter<number>();

  source: OTHubServerDataSource;

  ExportToJson() {
    const url = this.httpService.ApiUrl + '/api/nodes/dataholders?ercVersion=1&export=true&exporttype=0';
    window.location.href = url;
  }

  ExportToCsv() {
    const url = this.httpService.ApiUrl + '/api/nodes/dataholders?ercVersion=1&export=true&exporttype=1';
    window.location.href = url;
  }



  formatAmount(amount) {
    if (amount === null) {
      return null;
    }
    const split = amount.toString().split('.');
    let lastSplit = '';
    if (split.length === 2) {
      lastSplit = split[1];
      if (lastSplit.length > 3) {
        lastSplit = lastSplit.substring(0, 3);
      }

      if (lastSplit == '000')
      {
        return split[0];
      }

      return split[0] + '.' + lastSplit;
    }
    return split[0];
  }

  getIdentityIcon(identity: string) {
    return this.httpService.ApiUrl + '/api/icon/node/' + identity + '/' + (this.isDarkTheme ? 'dark' : 'light') + '/16';
  }

  pageSizeChanged(event) {
    this.source.setPaging(1, event, true);
  }

  // getNodes() {
  //   const headers = new HttpHeaders()
  //     .set('Content-Type', 'application/json')
  //     .set('Accept', 'application/json');
  //   let url = this.httpService.ApiUrl + '/api/nodes/dataholders?ercVersion=1';
  //   if (this.showOnlyMyNodes) {
  //     const myNodes = this.myNodeService.GetAll();
  //     // tslint:disable-next-line:prefer-for-of
  //     for (let index = 0; index < Object.keys(myNodes).length; index++) {
  //       const element = Object.values(myNodes)[index];
  //       url += '&identity=' + element.Identity;
  //     }
  //   } else if (this.managementWallet) {
  //     url += '&managementWallet=' + this.managementWallet;
  //   }
  //   url += '&' + (new Date()).getTime();
  //   return this.http.get<OTNodeSummaryModel[]>(url, { headers });
  // }

  ngOnDestroy() {
    // this.chRef.detach();

    // this.getNodesObserver.unsubscribe();
  }

  // Reload() {
  //   const self = this;

  //   const startTime = new Date();
  //   self.getNodesObserver = this.getNodes().subscribe(data => {
  //     const endTime = new Date();
  //     this.NodeModel = data;

  //     this.chRef.detectChanges();

  //     this.afterLoadWithCount.emit(this.NodeModel.length);

  //     if (!this.isTableInit) {
  //       this.isTableInit = true;
  //       const exportColumns = [0, 2, 3, 4, 5, 6, 7, 8];
  //       this.exportOptionsObj = {
  //         columns: exportColumns,
  //         format: {
  //           body(text, row, column, node) {
  //             return text;
  //           },
  //           header(text, column) {
  //             return text;
  //           }
  //         }
  //       };

  //       const table: any = $('.js-dataholders-table');
  //       this.dataTable = table.DataTable({
  //         responsive: true,
  //         pageLength: 25,
  //         columnDefs: [
  //           { targets: 0, visible: false },
  //           { targets: 1, visible: true },
  //           { targets: 2, visible: false },
  //           { targets: 3, visible: true },
  //           { targets: 4, visible: true },
  //           { targets: 5, visible: true },
  //           { targets: 6, visible: true },
  //           { targets: 7, visible: true }
  //         ],
  //         drawCallback() {
  //           $('img.lazy').lazyload();
  //         }
  //       });
  //     }

  //     const diff = endTime.getTime() - startTime.getTime();
  //     let minWait = 0;
  //     if (diff < 100) {
  //       minWait = 100 - diff;
  //     }
  //     setTimeout(() => {
  //       this.isLoading = false;
  //       if (this.NodeModel == null) {
  //         this.failedLoading = true;
  //       }
  //     }, minWait);
  //   }, err => {
  //     this.failedLoading = true;
  //     this.isLoading = false;
  //   });
  // }

  // copyToClipboard() {
  //   const that = { processing(isProcessing) { } };
  //   const e = null;
  //   const button = $.fn.dataTable.ext.buttons.copyHtml5;
  //   const config = this.exportOptionsObj;
  //   button.exportOptions = config;
  //   $.fn.dataTable.ext.buttons.copyHtml5.action.call(that, e, this.dataTable, config, button);
  // }

  // exportToCSV() {
  //   const that = { processing(isProcessing) { } };
  //   const e = null;
  //   const button = $.fn.dataTable.ext.buttons.csvHtml5;
  //   const config = this.exportOptionsObj;
  //   button.exportOptions = config;
  //   $.fn.dataTable.ext.buttons.csvHtml5.action.call(that, e, this.dataTable, config, button);
  // }


  // exportToExcel() {
  //   const that = { processing(isProcessing) { } };
  //   const e = null;
  //   const button = $.fn.dataTable.ext.buttons.excelHtml5;
  //   const config = this.exportOptionsObj;
  //   button.exportOptions = config;
  //   $.fn.dataTable.ext.buttons.excelHtml5.action.call(that, e, this.dataTable, config, button);
  // }

  // print() {
  //   const that = { processing(isProcessing) { } };
  //   const e = null;
  //   const button = $.fn.dataTable.ext.buttons.print;
  //   const config = this.exportOptionsObj;
  //   button.exportOptions = config;
  //   $.fn.dataTable.ext.buttons.print.action.call(that, e, this.dataTable, config, button);
  // }

  getNode(identity: string) {
    const headers = new HttpHeaders()
      .set('Content-Type', 'application/json')
      .set('Accept', 'application/json');
    const url = this.httpService.ApiUrl + '/api/nodes/dataholder/' + identity + '?' + (new Date()).getTime();
    return this.http.get<DataHolderDetailedModel>(url, { headers });
  }


  config: NbToastrConfig;
  toastStatus: NbComponentStatus;

  onEdit(event) {
    const oldData = event.data;
    const newData = event.newData;

    this.getNode(newData.Identity).subscribe(data => {
      if (data) {

        const oldModel = new MyNodeModel();
        oldModel.Identity = oldData.Identity;
        oldModel.DisplayName = '';
        this.myNodeService.Remove(oldModel);

        const model = new MyNodeModel();
        model.Identity = newData.Identity;
        model.DisplayName = newData.DisplayName;
        this.myNodeService.Add(model);
        this.resetSource();
        event.confirm.resolve();
        this.source.refresh();
      } else {
        this.config = new NbToastrConfig({duration: 8000});
        this.config.status = "warning";
        this.toastrService.show(
          'A node was not found by searching for the identity ' + newData.Identity + '. Please check you have entered the right identity.',  'Add Node', this.config);
      }
    });
  }

  onDelete(event) {
    var deleteData = event.data;

    const model = new MyNodeModel();
    model.Identity = deleteData.Identity;
    model.DisplayName = '';
    this.myNodeService.Remove(model);
    this.resetSource();
    event.confirm.resolve();
  }

  onCreate(event) {
  var newData = event.newData;

  this.getNode(newData.Identity).subscribe(data => {
    if (data) {
      const model = new MyNodeModel();
      model.Identity = newData.Identity;
      model.DisplayName = newData.DisplayName;
      this.myNodeService.Add(model);
      this.resetSource();
      event.confirm.resolve();
    } else {
      this.config = new NbToastrConfig({duration: 8000});
      this.config.status = "warning";
      this.toastrService.show(
        'A node was not found by searching for the identity ' + newData.Identity + '. Please check you have entered the right identity.',  'Add Node', this.config);
    }
  });
  }

  ngOnInit() {
    // this.isDarkTheme = $('body').hasClass('dark');
    // this.Reload();

    this.settings = {
      mode: 'inline',
      actions: {
        add: this.showOnlyMyNodes === 'true',
        edit: this.showOnlyMyNodes  === 'true',
        delete: this.showOnlyMyNodes  === 'true'
      },
      add: {
        addButtonContent: '<i class="nb-plus"></i>',
        createButtonContent: '<i class="nb-checkmark"></i>',
        cancelButtonContent: '<i class="nb-close"></i>',
        confirmCreate: true
      },
      edit: {
        editButtonContent: '<i class="nb-edit"></i>',
        saveButtonContent: '<i class="nb-checkmark"></i>',
        cancelButtonContent: '<i class="nb-close"></i>',
        confirmSave: true
      },
      delete: {
        deleteButtonContent: '<i class="nb-trash"></i>',
        confirmDelete: true,
      },
      columns: {
        LastSeenOnline: {
          title: 'Online',
          type: 'custom',
          renderComponent: OnlineIndicatorRenderComponent,
          filter: false,
          sort: false,
          editable: false,
          addable: false,
          // valuePrepareFunction: (value, row) => {
          //   return '<div style="font-size:30px;"><i class="nb-checkmark-circle"></i></div>';
          // }
        },
        Identity: {
          sort: false,
          title: 'Identity',
          type: 'html',
          filter: true,
          valuePrepareFunction: (value) => {
            if (!value) {
              return 'Unknown';
            }
  
            return '<a target=_self href="/nodes/dataholders/' + value +
             '""><img class="lazy" style="height:16px;width:16px;" title="' +
              value + '" src="' + this.getIdentityIcon(value) + '">' + value + '</a>';
          }
        },
        DisplayName: {
          title: 'Name',
          type: 'text',
          show: false,
          filter: false,
          sort: false,
          editable: true,
          addable: true,
        },
        TotalWonOffers: {
          sort: true,
          title: 'Jobs',
          type: 'number',
          filter: false,
          editable: false,
          addable: false
          // valuePrepareFunction: (value) => {
          //   return '<a class="navigateJqueryToAngular" href="/offers/' + value + '" onclick="return false;" title="' + value + '" >' + value.substring(0, 40) + '...</a>';
          // }
        },
        WonOffersLast7Days: {
          sort: true,
          sortDirection: 'desc',
          title: 'Jobs (7 Days)',
          type: 'number',
          filter: false,
          editable: false,
          addable: false
          // valuePrepareFunction: (value) => {
          //   const stillUtc = moment.utc(value).toDate();
          //   const local = moment(stillUtc).local().format('DD/MM/YYYY HH:mm');
          //   return local;
          // }
        },
        ActiveOffers: {
          sort: true,
          title: 'Active Jobs',
          type: 'number',
          filter: false,
          editable: false,
          addable: false
          // valuePrepareFunction: (value) => { return (value / 1000).toFixed(2).replace(/[.,]00$/, '') + ' KB';}
        },
        PaidTokens: {
          sort: true,
          title: 'Paidout Tokens',
          type: 'number',
          filter: false,
          editable: false,
          addable: false,
          valuePrepareFunction: (value) => {
            return this.formatAmount(value);
          }
        },
        StakeTokens: {
          sort: true,
          title: 'Staked Tokens',
          type: 'number',
          filter: false,
          editable: false,
          addable: false,
          valuePrepareFunction: (value) => {
            return this.formatAmount(value);
          }
        },
        StakeReservedTokens: {
          sort: true,
          title: 'Locked Tokens',
          type: 'number',
          filter: false,
          editable: false,
          addable: false,
          valuePrepareFunction: (value) => {
            return this.formatAmount(value);
          }
        }
      },
      pager: {
        display: true,
        perPage: 25
      }
    };

    if (this.showOnlyMyNodes !== 'true') {
      delete this.settings.columns.DisplayName;
      delete this.settings.columns.LastSeenOnline;
    }

    this.resetSource();
  }

  resetSource() {
    let url = this.httpService.ApiUrl + '/api/nodes/dataholders?ercVersion=1';
    if (this.showOnlyMyNodes === 'true') {
      const myNodes = this.myNodeService.GetAll();
      // tslint:disable-next-line:prefer-for-of
      const l = Object.keys(myNodes).length;
      for (let index = 0; index < l; index++) {
        const element = Object.values(myNodes)[index];
        url += '&identity=' + element.Identity;
      }

      if (l == 0) {
        url += "&identity=N/A";
      }
    } else if (this.managementWallet) {
      url += '&managementWallet=' + this.managementWallet;
    }

    if (this.source == null) {
    this.source = new OTHubServerDataSource(this.http, this.myNodeService,
      { endPoint: url });
    }
    else {
      this.source.ResetEndpoint(url);
    }
  }
}

class OTHubServerDataSource extends ServerDataSource {

  ResetEndpoint(endpoint: string) {
    this.conf.endPoint = endpoint;
  }

  constructor(http: HttpClient, private myNodeService: MyNodeService, conf?: ServerSourceConf | {}) {
    super(http, conf);
  }

  protected extractDataFromResponse(res: any): Array<any> {
    var data = super.extractDataFromResponse(res);

    data.forEach(element => {
      element.DisplayName = this.myNodeService.GetName(element.Identity, true);
    });
    return data;
  }

  public update(element, values): Promise<any> {
    return new Promise((resolve, reject) => {
        this.find(element).then(found => {
            //Copy the new values into element so we use the same instance
            //in the update call.
            // element.name = values.name;
            // element.enabled = values.enabled;
            // element.condition = values.condition;
            element.Identity = values.Identity;
            element.DisplayName = values.DisplayName;
            //Don't call super because that will cause problems - instead copy what DataSource.ts does.
            ///super.update(found, values).then(resolve).catch(reject);
            this.emitOnUpdated(element);
            this.emitOnChanged('update');
            resolve();
        }).catch(reject);
    });
}

  find(element) {
    const found = this.data.find(el => el.Identity == element.Identity);
    if (found) {
      return Promise.resolve(found);
    }
    return Promise.reject(new Error('Element was not found in the dataset'));
  }
}