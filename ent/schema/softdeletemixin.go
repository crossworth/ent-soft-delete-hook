package schema

import (
	"context"
	"time"

	entp "entgo.io/bug/ent"
	"entgo.io/bug/ent/hook"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type DeletedTimeAnnotation struct {
	OK bool
}

func (d DeletedTimeAnnotation) Name() string {
	return "DeletedTime"
}

type DeletedTime struct {
	mixin.Schema
}

func (DeletedTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_time").Optional(),
	}
}

func (DeletedTime) Annotations() []schema.Annotation {
	return []schema.Annotation{
		DeletedTimeAnnotation{
			OK: true,
		},
	}
}

type skipDeletedTimeHook struct{}

func WithSkipDeletedTimeHook(ctx context.Context) context.Context {
	return context.WithValue(ctx, skipDeletedTimeHook{}, true)
}

func (DeletedTime) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				_, skipDeletedTimeHook := ctx.Value(skipDeletedTimeHook{}).(bool)
				if skipDeletedTimeHook {
					return next.Mutate(ctx, m)
				}

				if idc, ok := m.(interface {
					IDs(ctx context.Context) ([]int, error)
					Client() *entp.Client
				}); ok {
					ids, err := idc.IDs(ctx)
					if err != nil {
						return nil, err
					}

					err = entp.SetDeletedTimeForType(ctx, idc.Client(), m.Type(), time.Now(), ids)
					if err != nil {
						return nil, err
					}

					return len(ids), nil
				}

				return next.Mutate(ctx, m)
			})
		}, ent.OpDeleteOne|ent.OpDelete,
		),
	}
}
