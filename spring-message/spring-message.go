/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package SpringMessage

import (
	"context"
)

// Message 消息体
type Message struct {
	Topic      string
	Key        string
	Value      []byte
	Properties map[string]string
}

// Consumer 消息消费者
type Consumer interface {
	Consume(ctx context.Context, msg *Message)
}

type SendArg struct {
	Properties map[string]string
}

type SendOption func(*SendArg)

// WithProperties 携带 properties 属性值
func WithProperties(properties map[string]string) SendOption {
	return func(arg *SendArg) {
		arg.Properties = properties
	}
}

// Producer 消息生产者
type Producer interface {
	Send(ctx context.Context, topic string, body string, options ...SendOption) error
}
