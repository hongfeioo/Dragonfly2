/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

type Application struct {
	Model
	Name              string             `gorm:"column:name;type:varchar(256);index:uk_application_name,unique;not null;comment:name" json:"name"`
	DownloadRateLimit uint               `gorm:"column:download_rate_limit;comment:download rate limit" json:"download_rate_limit"`
	URL               string             `gorm:"column:url;not null;comment:url" json:"url"`
	State             string             `gorm:"column:state;type:varchar(256);default:'enable';comment:state" json:"state"`
	BIO               string             `gorm:"column:bio;type:varchar(1024);comment:biography" json:"bio"`
	UserID            uint               `gorm:"comment:user id" json:"user_id"`
	User              User               `json:"-"`
	SchedulerClusters []SchedulerCluster `json:"-"`
	CDNClusters       []CDNCluster       `json:"-"`
}
