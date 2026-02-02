/*
 * Copyright 2021 Akamai Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { SelectableValue } from '@grafana/data';
import type { DataQuery, DataSourceJsonData } from '@grafana/schema';

export interface MyQuery extends DataQuery {
  selectedReport: SelectableValue<string>;
  zoneNames?: string;
  metricName?: string;
}

//export const defaultQuery: Partial<MyQuery> = {};

export const defaultQuery: Partial<MyQuery> = {
  selectedReport: { label: '', value: '' },
};


export interface MyDataSourceOptions extends DataSourceJsonData {
  clientSecret?: string;
  host?: string;
  accessToken?: string;
  clientToken?: string;
}
