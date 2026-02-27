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

import defaults from 'lodash/defaults';
import React, { ChangeEvent, PureComponent } from 'react';
import { QueryEditorProps, SelectableValue } from '@grafana/data';
import { InlineField, Input, Combobox, ComboboxOption } from '@grafana/ui';

import { DataSource } from './DataSource';
import { defaultQuery, MyDataSourceOptions, MyQuery } from './types';

type Props = QueryEditorProps<DataSource, MyQuery, MyDataSourceOptions>;

export class QueryEditor extends PureComponent<Props> {
  // loadReports function returns options asynchronously for Combobox
  getReports = async (input: string): Promise<ComboboxOption<string>[]> => {
    const uri = 'datasource/resource/openapireports';
    const results: SelectableValue<string>[] = await this.props.datasource.getResource(uri);

    return results
      .filter((item) => !input || item.label?.toLowerCase().includes(input.toLowerCase()))
      .map((item) => ({ label: item.label!, value: item.value! }));
  };

  onSelectReportsChange = (option: ComboboxOption<string> | null) => {
    const { onChange, query, onRunQuery } = this.props;

    if (option) {
      // Set the full SelectableValue for selectedReport
      const selectedReport: SelectableValue<string> = {
        label: option.label,
        value: option.value,
      };

      onChange({ ...query, selectedReport });

      if (query.zoneNames) {
        onRunQuery();
      }
    } else {
      // Option is null: user cleared selection, but since selectedReport is required, do nothing or keep old
    }
  };

  onZoneNamesChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, zoneNames: event.target.value });

    if (event.target.value && query.selectedReport) {
      onRunQuery();
    }
  };

  onMetricNameChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, metricName: event.target.value });

    if (query.zoneNames && query.selectedReport) {
      onRunQuery();
    }
  };

  render() {
    const query = defaults(this.props.query, defaultQuery);
    const { selectedReport, zoneNames, metricName } = query;

    return (
      <div className="gf-form">
        <div>
          <InlineField label="Report">
            <Combobox<string>
              placeholder="Select a report"
              options={this.getReports}
              value={selectedReport?.value}
              onChange={this.onSelectReportsChange}
              width="auto"
              minWidth={25}
            />
          </InlineField>

          <InlineField
            label="Zones"
            tooltip="Comma-separated zone names. Metrics for listed zones are added together."
          >
            <Input
              value={zoneNames || ''}
              onChange={this.onZoneNamesChange}
              placeholder="Enter zone names"
            />
          </InlineField>

          <InlineField
            label="Metric Name"
            tooltip="Graphed metric's name. If empty, a name is generated."
          >
            <Input
              value={metricName || ''}
              onChange={this.onMetricNameChange}
              placeholder="Enter metric name"
            />
          </InlineField>
        </div>
      </div>
    );
  }
}