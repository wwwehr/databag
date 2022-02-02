package databag

import (
  "databag/internal/store"
)

func getCardModel(card *store.Card) *Card {

  // populate group id list
  var groups []string;
  for _, group := range card.Groups {
    groups = append(groups, group.GroupId)
  }

  return &Card{
    CardId: card.CardId,
    NotifiedProfile: card.NotifiedProfile,
    NotifiedContent: card.NotifiedContent,
    NotifiedView: card.NotifiedView,
    CardProfile: &CardProfile{
      Guid: card.Guid,
      Handle: card.Username,
      Name: card.Name,
      Description: card.Description,
      Location: card.Location,
      Revision: card.ProfileRevision,
      ImageSet: card.Image != "",
      Version: card.Version,
      Node: card.Node,
    },
    CardData: &CardData {
      Revision: card.DataRevision,
      Status: card.Status,
      Notes: card.Notes,
      Token: card.OutToken,
      Groups: groups,
    },
  }
}

func getGroupModel(group *store.Group) *Group {
  return &Group{
    GroupId: group.GroupId,
    Revision: group.Revision,
    DataType: group.DataType,
    Data: group.GroupData.Data,
    Created: group.Created,
    Updated: group.Updated,
  }
}

func getArticleModel(article *store.Article, contact bool, shared bool) *Article {

  if !shared || article.ArticleData == nil {
    return &Article{
      ArticleId: article.ArticleId,
      Revision: article.Revision,
    }
  } else {

    var groups []string;
    if !contact {
      for _, group := range article.ArticleData.Groups {
        groups = append(groups, group.GroupId)
      }
    }

    var labels []string;
    for _, label := range article.ArticleData.Labels {
      labels = append(labels, label.LabelId)
    }

    return &Article{
      ArticleId: article.ArticleId,
      Revision: article.Revision,
      ArticleData: &ArticleData{
        DataType: article.ArticleData.DataType,
        Data: article.ArticleData.Data,
        Status: article.ArticleData.Status,
        Labels: labels,
        Groups: groups,
        TagCount: article.ArticleData.TagCount,
        Created: article.ArticleData.Created,
        Updated: article.ArticleData.Updated,
        TagUpdated: article.ArticleData.TagUpdated,
        TagRevision: article.ArticleData.TagRevision,
      },
    }
  }
}

