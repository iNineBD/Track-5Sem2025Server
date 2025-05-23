table "dim_project" {
  schema = schema.dw
  column "id" {
    null = false
    type = integer
    identity {
      generated = BY_DEFAULT
    }
  }
  column "name" {
    null = false
    type = character_varying(200)
  }
  column "description" {
    null = false
    type = character_varying(400)
  }
    column "created_date" {
    null = false
    type = timestamp
  }
    column "modified_date" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
}
table "dim_role" {
  schema = schema.dw
  column "id" {
    null = false
    type = integer
    identity {
      generated = BY_DEFAULT
    }
  }
  column "name" {
    null = false
    type = character_varying(200)
  }
  primary_key {
    columns = [column.id]
  }
}
table "dim_status" {
  schema = schema.dw
  column "id" {
    null = false
    type = integer
    identity {
      generated = BY_DEFAULT
    }
  }
  column "name" {
    null = false
    type = character_varying(200)
  }
  primary_key {
    columns = [column.id]
  }
}
table "dim_tag" {
  schema = schema.dw
  column "id" {
    null = false
    type = integer
    identity {
      generated = BY_DEFAULT
    }
  }
  column "name" {
    null = false
    type = character_varying(200)
  }
    column "color" {
    null = false
    type = character_varying(100)
  }
  primary_key {
    columns = [column.id]
  }
}
table "dim_user" {
  schema = schema.dw
  column "id" {
    null = false
    type = integer
    identity {
      generated = BY_DEFAULT
    }
  }
  column "full_name" {
    null = false
    type = character_varying(200)
  }
  column "color" {
    null = false
    type = character_varying(200)
  }
  column "id_role" {
    null = false
    type = integer
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "dim_user_id_role_fkey" {
    columns     = [column.id]
    ref_columns = [table.dim_role.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
}
table "fato_cards" {
  schema = schema.dw
  column "id_fato_card" {
    null = false
    type = integer
    identity {
      generated = BY_DEFAULT
    }
  }
  column "id_card" {
    null = false
    type = integer
  }
  column "id_project" {
    null = false
    type = integer
  }
  column "id_user" {
    null = false
    type = integer
  }
  column "id_status" {
    null = false
    type = integer
  }
  column "id_tag" {
    null = true
    type = integer
  }
  column "qtd_cards" {
    null    = false
    type    = integer
    default = 1
  }
  primary_key {
    columns = [column.id_fato_card]
  }
  foreign_key "fato_cards_id_project_fkey" {
    columns     = [column.id_project]
    ref_columns = [table.dim_project.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "fato_cards_id_status_fkey" {
    columns     = [column.id_status]
    ref_columns = [table.dim_status.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "fato_cards_id_tag_fkey" {
    columns     = [column.id_tag]
    ref_columns = [table.dim_tag.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "fato_cards_id_user_fkey" {
    columns     = [column.id_user]
    ref_columns = [table.dim_user.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
}
schema "dw" {
}
