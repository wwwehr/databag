/*
 * DataBag
 *
 * DataBag provides storage for decentralized identity based self-hosting apps. It is intended to support sharing of personal data and hosting group conversations. 
 *
 * API version: 0.0.1
 * Contact: roland.osborne@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package databag

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

  go SendNotifications()

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

  fs := http.FileServer(http.Dir("./"))
  router.PathPrefix("/static").Handler(http.StripPrefix("/", fs))

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AddAccount",
		strings.ToUpper("Post"),
		"/account/profile",
		AddAccount,
	},

	Route{
		"AddAccountApp",
		strings.ToUpper("Post"),
		"/account/apps",
		AddAccountApp,
	},

	Route{
		"AddAccountAuthentication",
		strings.ToUpper("Post"),
		"/account/auth",
		AddAccountAuthentication,
	},

	Route{
		"AddPublicAccount",
		strings.ToUpper("Post"),
		"/account/public/profile",
		AddPublicAccount,
	},

	Route{
		"GetAccountApps",
		strings.ToUpper("Get"),
		"/account/apps",
		GetAccountApps,
	},

	Route{
		"GetAccountAsset",
		strings.ToUpper("Get"),
		"/account/assets/{assetId}",
		GetAccountAsset,
	},

	Route{
		"GetAccountDid",
		strings.ToUpper("Get"),
		"/account/did",
		GetAccountDid,
	},

	Route{
		"GetAccountImage",
		strings.ToUpper("Get"),
		"/account/profile/image",
		GetAccountImage,
	},

	Route{
		"GetAccountProfile",
		strings.ToUpper("Get"),
		"/account/profile",
		GetAccountProfile,
	},

	Route{
		"GetAccountStatus",
		strings.ToUpper("Get"),
		"/account/status",
		GetAccountStatus,
	},

	Route{
		"GetAccountToken",
		strings.ToUpper("Get"),
		"/account/token",
		GetAccountToken,
	},

	Route{
		"GetAccountUsername",
		strings.ToUpper("Get"),
		"/account/available",
		GetAccountUsername,
	},

	Route{
		"GetPublicStatus",
		strings.ToUpper("Get"),
		"/account/public/status",
		GetPublicStatus,
	},

	Route{
		"RemoveAccount",
		strings.ToUpper("Delete"),
		"/account/profile",
		RemoveAccount,
	},

	Route{
		"RemoveAccountApp",
		strings.ToUpper("Delete"),
		"/account/apps/{appId}",
		RemoveAccountApp,
	},

	Route{
		"SetAccountApp",
		strings.ToUpper("Put"),
		"/account/apps",
		SetAccountApp,
	},

	Route{
		"SetAccountAuthentication",
		strings.ToUpper("Put"),
		"/account/auth",
		SetAccountAuthentication,
	},

	Route{
		"SetAccountExport",
		strings.ToUpper("Put"),
		"/account/export",
		SetAccountExport,
	},

	Route{
		"SetAccountNode",
		strings.ToUpper("Put"),
		"/account/node",
		SetAccountNode,
	},

	Route{
		"AddNodeAccount",
		strings.ToUpper("Post"),
		"/admin/accounts",
		AddNodeAccount,
	},

	Route{
		"GetNodeAccountImage",
		strings.ToUpper("Get"),
		"/admin/accounts/{accountId}/image",
		GetNodeAccountImage,
	},

	Route{
		"GetNodeAccounts",
		strings.ToUpper("Get"),
		"/admin/accounts",
		GetNodeAccounts,
	},

	Route{
		"GetNodeConfig",
		strings.ToUpper("Get"),
		"/admin/config",
		GetNodeConfig,
	},

	Route{
		"GetNodeStatus",
		strings.ToUpper("Get"),
		"/admin/status",
		GetNodeStatus,
	},

	Route{
		"ImportAccount",
		strings.ToUpper("Post"),
		"/admin/accounts/import",
		ImportAccount,
	},

	Route{
		"RemoveNodeAccount",
		strings.ToUpper("Delete"),
		"/admin/accounts/{accountId}",
		RemoveNodeAccount,
	},

	Route{
		"SetNodeAccount",
		strings.ToUpper("Put"),
		"/admin/accounts/{accountId}/reset",
		SetNodeAccount,
	},

	Route{
		"SetNodeConfig",
		strings.ToUpper("Put"),
		"/admin/config",
		SetNodeConfig,
	},

	Route{
		"SetNodeStatus",
		strings.ToUpper("Put"),
		"/admin/status",
		SetNodeStatus,
	},

	Route{
		"AddGroup",
		strings.ToUpper("Post"),
		"/alias/groups",
		AddGroup,
	},

	Route{
		"GetGroups",
		strings.ToUpper("Get"),
		"/alias/groups",
		GetGroups,
	},

	Route{
		"RemoveGroup",
		strings.ToUpper("Delete"),
		"/alias/groups/{groupId}",
		RemoveGroup,
	},

	Route{
		"SetGroupSubject",
		strings.ToUpper("Put"),
		"/alias/groups/{groupId}/subject",
		SetGroupSubject,
	},

	Route{
		"AddArticle",
		strings.ToUpper("Post"),
		"/attribute/articles",
		AddArticle,
	},

	Route{
		"ClearArticleGroup",
		strings.ToUpper("Delete"),
		"/attribute/articles/{articleId}/groups/{groupId}",
		ClearArticleGroup,
	},

	Route{
		"GetArticleSubjectField",
		strings.ToUpper("Get"),
		"/attribute/articles/{articleId}/subject/{field}",
		GetArticleSubjectField,
	},

	Route{
		"GetArticles",
		strings.ToUpper("Get"),
		"/attribute/articles",
		GetArticles,
	},

	Route{
		"RemoveArticle",
		strings.ToUpper("Delete"),
		"/attribute/articles/{articleId}",
		RemoveArticle,
	},

	Route{
		"SetArticleGroup",
		strings.ToUpper("Put"),
		"/attribute/articles/{articleId}/groups/{groupId}",
		SetArticleGroup,
	},

	Route{
		"SetArticleSubject",
		strings.ToUpper("Put"),
		"/attribute/articles/{articleId}/subject",
		SetArticleSubject,
	},

	Route{
		"Authorize",
		strings.ToUpper("Put"),
		"/authorize",
		Authorize,
	},

	Route{
		"AddCard",
		strings.ToUpper("Post"),
		"/contact/cards",
		AddCard,
	},

	Route{
		"ClearCardGroup",
		strings.ToUpper("Delete"),
		"/contact/cards/{cardId}/groups/{groupId}",
		ClearCardGroup,
	},

	Route{
		"ClearCardNotes",
		strings.ToUpper("Delete"),
		"/contact/cards/{cardId}/notes",
		ClearCardNotes,
	},

	Route{
		"GetCardDetail",
		strings.ToUpper("Get"),
		"/contact/cards/{cardId}/detail",
		GetCardDetail,
	},

	Route{
		"GetCardProfile",
		strings.ToUpper("Get"),
		"/contact/cards/{cardId}/profile",
		GetCardProfile,
	},

	Route{
		"GetCardProfileImage",
		strings.ToUpper("Get"),
		"/contact/cards/{cardId}/profile/image",
		GetCardProfileImage,
	},

	Route{
		"GetCards",
		strings.ToUpper("Get"),
		"/contact/cards",
		GetCards,
	},

	Route{
		"GetCloseMessage",
		strings.ToUpper("Get"),
		"/contact/cards/{cardId}/closeMessage",
		GetCloseMessage,
	},

	Route{
		"GetOpenMessage",
		strings.ToUpper("Get"),
		"/contact/cards/{cardId}/openMessage",
		GetOpenMessage,
	},

	Route{
		"RemoveCard",
		strings.ToUpper("Delete"),
		"/contact/cards/{cardId}",
		RemoveCard,
	},

	Route{
		"SetArticleRevision",
		strings.ToUpper("Put"),
		"/contact/article/revision",
		SetArticleRevision,
	},

	Route{
		"SetCardGroup",
		strings.ToUpper("Put"),
		"/contact/cards/{cardId}/groups/{groupId}",
		SetCardGroup,
	},

	Route{
		"SetCardNotes",
		strings.ToUpper("Put"),
		"/contact/cards/{cardId}/notes",
		SetCardNotes,
	},

	Route{
		"SetCardProfile",
		strings.ToUpper("Put"),
		"/contact/cards/{cardId}/profile",
		SetCardProfile,
	},

	Route{
		"SetCardStatus",
		strings.ToUpper("Put"),
		"/contact/cards/{cardId}/status",
		SetCardStatus,
	},

	Route{
		"SetChannelRevision",
		strings.ToUpper("Put"),
		"/contact/channel/revision",
		SetChannelRevision,
	},

	Route{
		"SetCloseMessage",
		strings.ToUpper("Put"),
		"/contact/closeMessage",
		SetCloseMessage,
	},

	Route{
		"SetOpenMessage",
		strings.ToUpper("Put"),
		"/contact/openMessage",
		SetOpenMessage,
	},

	Route{
		"SetProfileRevision",
		strings.ToUpper("Put"),
		"/contact/profile/revision",
		SetProfileRevision,
	},

	Route{
		"SetViewRevision",
		strings.ToUpper("Put"),
		"/contact/view/revision",
		SetViewRevision,
	},

	Route{
		"AddChannel",
		strings.ToUpper("Post"),
		"/content/channels",
		AddChannel,
	},

	Route{
		"AddChannelTopicAsset",
		strings.ToUpper("Post"),
		"/content/channels/{channelId}/topics/{topicId}/assets",
		AddChannelTopicAsset,
	},

	Route{
		"AddChannelTopic",
		strings.ToUpper("Post"),
		"/content/channels/{channelId}/topics",
		AddChannelTopic,
	},

	Route{
		"AddChannelTopicTag",
		strings.ToUpper("Post"),
		"/content/channels/{channelId}/topics/{topicId}/tags",
		AddChannelTopicTag,
	},

	Route{
		"ClearChannelCard",
		strings.ToUpper("Delete"),
		"/content/channels/{channelId}/cards/{cardId}",
		ClearChannelCard,
	},

	Route{
		"ClearChannelGroup",
		strings.ToUpper("Delete"),
		"/content/channels/{channelId}/groups/{groupId}",
		ClearChannelGroup,
	},

	Route{
		"GetChannelAsset",
		strings.ToUpper("Get"),
		"/content/channels/{channelId}/topics/{topicId}/assets/{assetId}",
		GetChannelAsset,
	},

	Route{
		"GetChannelAssets",
		strings.ToUpper("Get"),
		"/content/channels/{channelId}/topics/{topicId}/assets",
		GetChannelAssets,
	},

	Route{
		"GetChannel",
		strings.ToUpper("Get"),
		"/content/channels/{channelId}",
		GetChannel,
	},

	Route{
		"GetChannelSubjectField",
		strings.ToUpper("Get"),
		"/content/channels/{channelId}/subject/{field}",
		GetChannelSubjectField,
	},

	Route{
		"GetChannelTopic",
		strings.ToUpper("Get"),
		"/content/channels/{channelId}/topics/{topicId}",
		GetChannelTopic,
	},

	Route{
		"GetChannelTopicDetail",
		strings.ToUpper("Get"),
		"/content/channels/{channelId}/topics/{topicId}/detail",
		GetChannelTopicDetail,
	},

	Route{
		"GetChannelTopicCount",
		strings.ToUpper("Get"),
		"/content/channels/{channelId}/topics/{topicId}/count",
		GetChannelTopicCount,
	},

	Route{
		"GetChannelTopicSubjectField",
		strings.ToUpper("Get"),
		"/content/channels/{channelId}/topics/{topicId}/subject/{field}",
		GetChannelTopicSubjectField,
	},

	Route{
		"GetChannelTopicTagSubjectField",
		strings.ToUpper("Get"),
		"/content/channels/{channelId}/topics/{topicId}/tags/{tagId}/subject/{field}",
		GetChannelTopicTagSubjectField,
	},

	Route{
		"GetChannelTopicTags",
		strings.ToUpper("Get"),
		"/content/channels/{channelId}/topics/{topicId}/tags",
		GetChannelTopicTags,
	},

	Route{
		"GetChannelTopics",
		strings.ToUpper("Get"),
		"/content/channels/{channelId}/topics",
		GetChannelTopics,
	},

	Route{
		"GetChannels",
		strings.ToUpper("Get"),
		"/content/channels",
		GetChannels,
	},

	Route{
		"RemoveChannel",
		strings.ToUpper("Delete"),
		"/content/channels/{channelId}",
		RemoveChannel,
	},

	Route{
		"RemoveChannelAsset",
		strings.ToUpper("Delete"),
		"/content/channels/{channelId}/topics/{topicId}/assets/{assetId}",
		RemoveChannelAsset,
	},

	Route{
		"RemoveChannelTopic",
		strings.ToUpper("Delete"),
		"/content/channels/{channelId}/topics/{topicId}",
		RemoveChannelTopic,
	},

	Route{
		"RemoveChannelTopicTag",
		strings.ToUpper("Delete"),
		"/content/channels/{channelId}/topics/{topicId}/tags/{tagId}",
		RemoveChannelTopicTag,
	},

	Route{
		"SetChannelCard",
		strings.ToUpper("Put"),
		"/content/channels/{channelId}/cards/{cardId}",
		SetChannelCard,
	},

	Route{
		"SetChannelTopicConfirmed",
		strings.ToUpper("Put"),
		"/content/channels/{channelId}/topics/{topicId}/confirmed",
		SetChannelTopicConfirmed,
	},

	Route{
		"SetChannelGroup",
		strings.ToUpper("Put"),
		"/content/channels/{channelId}/groups/{groupId}",
		SetChannelGroup,
	},

	Route{
		"SetChannelSubject",
		strings.ToUpper("Put"),
		"/content/channels/{channelId}/subject",
		SetChannelSubject,
	},

	Route{
		"SetChannelTopicSubject",
		strings.ToUpper("Put"),
		"/content/channels/{channelId}/topics/{topicId}/subject",
		SetChannelTopicSubject,
	},

	Route{
		"SetChannelTopicTagSubject",
		strings.ToUpper("Put"),
		"/content/channels/{channelId}/topics/{topicId}/tags/{tagId}/subject",
		SetChannelTopicTagSubject,
	},

	Route{
		"GetProfile",
		strings.ToUpper("Get"),
		"/profile",
		GetProfile,
	},

	Route{
		"GetProfileImage",
		strings.ToUpper("Get"),
		"/profile/image",
		GetProfileImage,
	},

	Route{
		"GetProfileMessage",
		strings.ToUpper("Get"),
		"/profile/message",
		GetProfileMessage,
	},

	Route{
		"SetProfile",
		strings.ToUpper("Put"),
		"/profile/data",
		SetProfile,
	},

	Route{
		"SetProfileImage",
		strings.ToUpper("Put"),
		"/profile/image",
		SetProfileImage,
	},

	Route{
		"Status",
		strings.ToUpper("Get"),
		"/status",
		Status,
	},
}

